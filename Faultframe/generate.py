import re
from re import sub
import os
import math
import shutil
import random

class codeGenerate:
    def __init__(self, file_path, template_path):
        self.Name = ''  # 设备名称
        self.StateVarible=[]  #所有故障模型的状态变量列表
        self.Evidence_list=[]#所有观测变量的列表
        self.Transfer=[] #转移模型初始概率
        self.file_path = file_path    #用户输入文件路径
        self.templatepath = template_path  # 模板路径
        self.State_Input = []  # 状态变量输入接口
        self.State_Input_var = []  # 状态变量输入
        self.Sensor = []  #传感器模型公式
        self.newfile = shutil.copyfile(self.templatepath, self.Name + '_Code.go')
        '''
        self.alter('className', self.Name)
        name = self.Name.lower()
        self.name = name
        self.alter('name', self.name)
        slice = self.Name + 's'
        self.alter('slice', slice)
        sheetName = self.Name[:-5]
        self.alter("sheetName", sheetName)
        typefaults = self.name + 's'
        self.alter('typefaults', typefaults)
        '''


    def dataRead(self):
        file = open(self.file_path, 'r', encoding='utf-8')
        file_lines = file.readlines()
        file.close()#未完
        l = 0  # 行计数
        for line in file_lines:  # 遍历文件所有行
            l += 1
            if line.find('Name') != -1:
                self.Name = line.split('=')[1].strip()  # 获取设备名字

            if line.find('StateVarible') != -1:
                for item in file_lines[l:]:
                    if item.find('}') != -1:
                        break
                    item = item.strip('\n').strip(',').replace(' ', '')
                    item_list = item.split(',')
                    self.StateVarible+= item_list  # 获取状态量

            if line.find('Evidence') != -1:
                for item in file_lines[l:]:
                    if item.find('}') != -1:
                        break
                    item = item.strip('\n').strip(',').replace(' ', '')
                    item_list = item.split(',')
                    self.Evidence_list+= item_list  # 获取观测量

            if line.find('Transfer')!= -1:
                item_list = []
                for item in file_lines[l:]:
                    if item.find('}') != -1:
                        break
                    num1= re.finditer('-?\d+(\.\d+)?',item)
                    for match in num1:
                        item_list.append(float(match.group()))
                self.Transfer =item_list
                print('self.Transfer',self.Transfer)

            self.Equation_Get(line, 'Sensor', file_lines, l, self.Sensor)
        print('self.sensor',self.Sensor)


    def List_to_Str(self,List):     #将列表中的元素转化为字符串
        Str=''
        for item in List:
            if item ==List[-1]:
                Str+=item
            else:
                Str+=item+','
        return Str

    def Equation_Get(self,line,Equation_name,file_lines,l,Equation_list):  #此函数用于获取元件的各种公式，后续反复调用
        if line.find(Equation_name)!=-1:
            for item in file_lines[l:]:
                if item.find('}')!=-1:
                        break
                item=item.strip('\n').replace(' ','')
                Equation_list.append(item)

    def insert(self,initial,begin,end,supplement,insert_file):
        s = initial
        l = 0
        with open(self.Name + "_Code.go", "r+", encoding="utf-8") as f_tem:
            file_lines = f_tem.readlines()
            for line in file_lines:
                l += 1
                if line.find(begin) != -1:
                    for item in file_lines[l:]:
                        if item.find(end) != -1:
                            break
                        s += str(item)  # 获取Generate
        s += supplement
        with open(insert_file, "r+", encoding="utf-8") as f:
            ls = f.read()
            f.write(s)

    def insert_part(self,file,text,nextrow):
        t=text
        f=file
        with open(f, "r", encoding="utf-8") as f_tem:
            a = f_tem.readlines()
            for i in range(len(a)):
                if a[i].find(nextrow) != -1:
                    a.insert(i, t)
                    break
        with open(f, "w", encoding="utf-8") as f_f:
            for i in a:
                f_f.write(i)
            f_f.close()

    def structGenerate(self):
        self.newfile = shutil.copyfile(self.templatepath, self.Name + '_Code.go')
        # 复制模板，生成新的元件代码文件，文件名为‘元件_Code’
        self.alter('className', self.Name)
        self.alter('States', self.List_to_Str(self.StateVarible))

    def alter(self, old_str, new_str):
        newfile = self.newfile
        with open(newfile, "r", encoding="utf-8") as f1, open("%s.bak" % newfile, "w", encoding="utf-8") as f2:
            for line in f1:
                f2.write(sub('%\(' + old_str + '\)', new_str, line))
        os.remove(newfile)
        os.rename("%s.bak" % newfile, newfile)

    def typeGenerate(self):
        self.alter('Evidence_list', self.List_to_Str(self.Evidence_list))
        self.alter('LineState', self.List_to_Str(self.StateVarible))
        name=self.Name.lower()
        self.name = name
        self.alter('name',self.name)
        s=''
        s+='package faultmodelstruct\n\n'+'type '+self.Name+' struct {\n'
        l = 0
        with open(self.Name+"_Code.go", "r+", encoding="utf-8") as f_tem:
            file_lines=f_tem.readlines()
            for line in file_lines:
                l+=1
                if line.find('type '+self.Name+' struct') != -1:
                    for item in file_lines[l:]:
                        if item.find('}') != -1:
                            break
                        s += str(item)  # 获取typeGenerate
        s+='}\n'
        f=open('faultmodelstruct/'+self.name+'.go','a+')
        #ls=f.read()
        f.write(s)

    def insert_basefault(self):
        slice = self.Name + 's'
        textt='\t'+slice +'[]'+self.Name+'\n'
        self.insert_part("faultmodelstruct/basefault.go", textt, "//插入故障列表")
        text ='\t'+ self.Name+'Names map[string]int32\n'
        self.insert_part("faultmodelstruct/basefault.go",text,"//故障名与切片下标")

    def addGenerate(self):
    #将函数xxxFaultAdd嵌入add.go中
        slice=self.Name+'s'
        self.alter('slice',slice)
        sheetName = self.Name[:-5]
        self.alter("sheetName",sheetName)
        self.insert('func '+self.Name+'Add(xlsx *excelize.File) []faultmodelstruct.'+self.Name+'{\n',"func "+self.Name + "Add",'return '+slice,"\treturn "+slice+"\n"+"}\n\n","baseFunction/add.go")
        text='\tfault.'+slice+ '= '+self.Name+'Add(xlsx)\n'
        self.insert_part("baseFunction/add.go",text,"//调用FaultAdd函数")


    def SetEvidenceGenerate(self):
        Ny_code=''
        name = self.Name.lower()
        sheetName = self.Name[:-5]
        self.name = name
        typefaults = self.name + 's'
        self.alter('typefaults', typefaults)
        self.alter('sheetName', sheetName)
        for evidence in self.Evidence_list:  # 遍历观测变量列表
            Ny_code += '\t' + '\t'+'common.Evidences[' + typefaults + '[i].'+evidence +'] = evidence.'+evidence+'\n'
            # 给观测变量序列赋值,如 common.Evidences[linefaults[i].I] = evidence.I
        self.alter('Y_InputCode',Ny_code)
        self.insert("func Set" + sheetName + "Evidence(faults *faultmodelstruct.Fault,common *commonStruct.Common,evidence *commonStruct.Evidence){\n",
        "func Set"+sheetName + "Evidence",'//给观测变量序列赋值',"\t//给观测变量序列赋值\n\t}\n}\n\n","baseFunction/setevidence.go")



#className=LineFault name=linefault sheetName=Line  %slice=LineFaults %(typefaults)=linefaults  self.Name= LineFault
    def IndexGenerate(self):
        className=self.Name
        name = self.Name.lower()
        sheetName=self.Name[:-5]
        slice=self.Name+'s'
        self.name=name
        typefaults=self.name+'s'

        self.alter('name', name)
        self.alter('className', className)
        self.alter('sheetName', sheetName)
        self.alter('slice', slice)
        self.alter('sheetName', sheetName)
        self.alter('typefaults',typefaults)

        state_count = 0  # 状态变量计数
        Ny_count = 0  #观测变量计数
        state_code = ''  # 状态变量映射的代码
        Ny_code = ''  # 代数变量映射的代码

        for state in self.StateVarible:  #遍历状态变量列表
            state=name[:]+'.'+state  #状态变量改变形式,如LineState 变成 linefault.LineState
            if state_count==0:
                state_code+='\t'+'\t'+state+' = '+'common.Nx'+'\n'  #如果状态变量计数为0，此处相当于+0，故省去+，如：linefault.LineState = Common.Nx
            else:
                state_code+='\t'+'\t'+state+' = '+'common.Nx + '+str(state_count)+'\n'  #状态变量映射到Nx,如：Common.Nx += 1
            state_count+=1  #状态变量计数+1

        for evidence in self.Evidence_list:  #遍历观测变量列表
            evidence=name[:]+'.'+evidence  #代数变量改变形式, 如Load变为linefault.Load
            if Ny_count==0:
                Ny_code+='\t'+'\t'+evidence+' = '+'common.Ny'+'\n'  #如果代数变量计数为0，此处相当于+0，故省去+，如linefault.Load = Common.Ny
            else:
                Ny_code+='\t'+'\t'+evidence+' = '+'common.Ny + '+str(Ny_count)+'\n'  #代数变量映射到Ny,如：linefault.Temperature = Common.Ny + 1
            Ny_count+=1  #观测变量计数+1

        state_code += '\t' + '\t' + 'common.Nx += ' + str(state_count) + '\n'  # Common.Nx+=最终状态变量计数
        Ny_code += '\t' + '\t' + 'common.Ny += ' + str(Ny_count) + '\n'  # Common.Ny+=最终代数变量计数
        self.alter('X_IndexCode', state_code)
        self.alter('Y_IndexCode', Ny_code)
        self.insert("func " + className + "Index(fault *faultmodelstruct.Fault,common *commonStruct.Common){\n",
                    "func " + className + "Index",'//观测变量映射', "\t//观测变量映射\n\t}\n}\n\n","baseFunction/index.go")
        text='\t'+self.Name+'Index(fault,common)\n'
        self.insert_part("baseFunction/index.go",text,'//添加FaultIndex函数')

    def SensorGenerate(self):
        P_sensor_state=''
        P_sensor=''
        i=0
        for evidence in self.Evidence_list:
            i+=1
            P_sensor_state +=  '\t' + 'P'+ evidence +','
            if i<len(self.Evidence_list):
                P_sensor += '\t' + 'P' + evidence + ','
            else:
                P_sensor += '\t' + 'P' + evidence
        P_sensor_state += 'float64'
        sheetName = self.Name[:-5]
        self.alter('sheetName',sheetName)
        self.alter('P_sensor_state',P_sensor_state)
        self.alter('cal_Psensor', self.Name)
        self.alter('P_sensor', P_sensor)
        self.insert("func " + sheetName + "SensorModel(sensor []float64, GaussCoefficient [][]float64) [2]float64 {\n",
                    "func " + sheetName + "SensorModel",'//计算传感器模型概率',"\t//计算传感器模型概率\n}\n\n","baseFunction/sensor.go")

    def TransferGenerate(self):
        P_trans_state = ''
        for state in self.StateVarible:
            P_trans_state += '\t' + 'P' + state + ','
        P_trans_state += 'float64'
        #for p_transs in self.Transfer:
        #未完成，需要将用户输入的初始概率计算出来，再存入common.Trans 和Pensor_t中
        sheetName = self.Name[:-5]
        self.alter('sheetName', sheetName)
        self.alter('P_trans_state', P_trans_state)
        self.insert(
            "func " + sheetName + "TransModel(transP float64,index int,common *commonStruct.Common,faults *faultmodelstruct.Fault) float64 {\n",
            "func " + sheetName + "TransModel", '//计算完善的转移模型概率', "\t//计算完善的转移模型概率\n}\n\n", "baseFunction/transfer.go")
        P0 = str(self.Transfer[0])
        #P1 = self.Transfer[1]
        P2 = str(self.Transfer[2])
        #P3 = self.Transfer[3]
        self.insert_part("baseFunction/transfer.go","\t\tP_sensor ="+P2,"//插入"+sheetName+"P2")
        self.insert_part("baseFunction/transfer.go", "\t\tP_sensor=" + P0, "//插入" + sheetName + "P0")




    def calculatePGenerate(self):
        sheetName = self.Name[:-5]
        self.alter('sheetName', sheetName)
        self.insert("func Calculate" + sheetName + "P(observation [][]float64, beliefP float64, transP float64, GaussCoefficient [][]float64,index int,faults *faultmodelstruct.Fault,common *commonStruct.Common) ([]float64,[]bool) {\n",
                    "func Calculate" + sheetName+"P", '//返回某单个元件的时序故障概率', "\t//返回某单个元件的时序故障概率\n}\n\n", "baseFunction/calculateP.go")



    def initialGenerate(self):
        typefaults = self.name + 's'
        name=self.name
        slice=self.Name+'s'
        text='\tfor i,_ := range('+typefaults+'){\n'+'\t\t'+name +':='+typefaults+'[i]\n'
        i=0
        q=0
        sheetName = self.Name[:-5]
        for evidence in self.Evidence_list:
            if i<len(self.Evidence_list):
                text += '\t\tcommon.Evidences['+name+'.'+evidence+'] = 0\n'
            i += 1
        for state in self.StateVarible:
            if q < len(self.StateVarible):
                text+='\t\tcommon.States['+name+'.'+ state+ '] = false\n'
            q += 1
        text+='\t}\n'
        self.insert_part("baseFunction/initial.go",text,"//加入变量成功")
        textdef='\t'+ typefaults +':= fault.'+slice+'\n'
        self.insert_part("baseFunction/initial.go",textdef,"//添加元件故障序列")


    def creat_fault_model(self):
        self.dataRead()
        self.structGenerate()
        self.typeGenerate()
        self.insert_basefault()
        self.addGenerate()
        self.IndexGenerate()
        self.SetEvidenceGenerate()
        self.SensorGenerate()
        self.TransferGenerate()
        self.calculatePGenerate()
        self.initialGenerate()


if __name__=='__main__':
    code_line=codeGenerate('faultmodels/linefault', 'template.txt')
    code_line.creat_fault_model()
    code_trans = codeGenerate('faultmodels/transfault', 'template.txt')
    code_trans.creat_fault_model()

