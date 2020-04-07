import os
from index import *
import re


def construct_name(slug):
    res = ""
    temp_list = slug.split("-")
    for word in temp_list:
        res += word.capitalize()
    return res


def strong(str):
    return "**"+str+"**"

tagdict = {}
orderdict = {}

link = "https://github.com/LZH139/leetcode_Go/blob/master/note/HashTable/simple/"
for root, dirs, files in os.walk(os.getcwd()+"/note"):
    for file in files:
        path = os.path.join(root, file)
        if path.split(".")[-1] == "md":
            with open(os.path.join(root, file), "r", encoding='utf-8') as f:
                content = f.read()

                if content.find("```go\n\n\n```") == -1:

                    # 更新笔记里的项目文件链接
                    templist = path.split("/")

                    tag = templist[-3]
                    difficulty = templist[-2]
                    noteFileName = templist[-1]
                    topicOrder = noteFileName.split(".")[0]
                    codeFileNme = construct_name(index[noteFileName.split(".")[-2].strip()])

                    # 面试题的序列包含中文字符，直接让其排到最后
                    if not topicOrder.isdigit():
                        topicOrder = 100000

                    noteFileLink = link + noteFileName.replace(".", "%2E").replace(" ", "%20")
                    codeFileLink = link.replace("note", "src") + codeFileNme + "/" + codeFileNme+".go"

                    replaceplace = re.findall("\[代码文件\]\(.*\)", content)[0]
                    content = content.replace(replaceplace, "[代码文件](" + codeFileLink + ")")
                    with open(os.path.join(root, file), "w+") as f:
                        f.write(content)

                    if tag in tagdict:
                        tagdict[tag][int(topicOrder)] = [noteFileName, difficulty, codeFileLink]
                    else:
                        tagdict[tag] = {int(topicOrder): [noteFileName, difficulty, codeFileLink]}
                    orderdict[int(topicOrder)] = [noteFileName, difficulty, codeFileLink]

tagtable = ""
ordertable = ""

for key in tagdict.keys():
    tagtable += "#### "+key + "\n" + head + layout
    for v in sorted(tagdict[key]):
        imfList = tagdict[key][v]
        tagtable += "| "+strong(imfList[0])+" | "+strong(imfList[1])+" | "+strong("[GO]("+imfList[2].replace(".", "%2E").replace(" ", "%20")+")")+" |\n"
    tagtable+="\n"

ordertable +="#### Order list"+ "\n" + head + layout
for v in sorted(orderdict):
    imfList = orderdict[v]
    ordertable += "| "+strong(imfList[0])+" | "+strong(imfList[1])+" | "+strong("[GO]("+imfList[2].replace(".", "%2E").replace(" ", "%20")+")")+" |\n"
ordertable += "\n"

with open("../leetcode_Go/README.md", "r") as f:
    content = f.read()

content = content.split("Portals")
content = content[0] + "Portals >>>\n" + tagtable + ordertable + "## <<< Portals" + content[2]

with open("../leetcode_Go/README.md", "w+") as f:
    f.write(content)

print("Portals >>>\n" + tagtable + ordertable + "## <<< Portals")