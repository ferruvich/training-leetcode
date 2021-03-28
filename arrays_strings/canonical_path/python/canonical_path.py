class Solution:
    def simplify_path(self, path: str) -> str:
        pathElements = []
        
        contents = path.split("/")
        for el in contents:
            if not el or el == ".":
                continue
            elif el=="..":
                if pathElements:
                    pathElements.pop()   
            else:
                pathElements.append(el)
        
        return "/"+"/".join(pathElements)