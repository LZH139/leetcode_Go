import os
from index import *


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

link = "https://github.com/LZH139/leetcode_Go/blob/master/note"
for root, dirs, files in os.walk(os.getcwd()+"/note"):
    for file in files:
        path = os.path.join(root, file)
        if path.split(".")[-1] == "md":
            f = open(os.path.join(root, file), "r", encoding='utf-8')
            content = f.read()
            f.close()
            if content.find("```go\n\n\n```") == -1:

                # 更新笔记里的项目文件
                templist = path.split("/")

                tag = templist[-3]
                difficulty = templist[-2]
                topicName = templist[-1]
                topicOrder = topicName.split(".")[0]

                githublink = link + "/" + tag + "/" + difficulty + "/" + topicName

                content = content.replace("[代码文件]()", "[代码文件]("+githublink.replace(".", "%2E").replace(" ", "%20")+")")
                f = open(os.path.join(root, file), "w+")
                f.write(content)
                f.close()

                if tag in tagdict:
                    tagdict[tag][int(topicOrder)] = [topicName, difficulty, githublink]
                else:
                    tagdict[tag] = {int(topicOrder): [topicName, difficulty, githublink]}
                orderdict[int(topicOrder)] = [topicName, difficulty, githublink]

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

with open("README.md", "r") as f:
    content = f.read()

content = content.split("Portals")
content = content[0] + "Portals >>>\n" + tagtable + ordertable + "## <<< Portals" + content[2]

with open("README.md", "w+") as f:
    f.write(content)