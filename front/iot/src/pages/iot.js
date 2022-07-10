import * as React from "react"
const getPages= async()=>{
    const res=await fetch("https://raw.githubusercontent.com/Tegei/tegei.github.io/main/memberList.csv")
    const text=await res.text()
    return text
}
const IndexPage = () => {
    const [devices,setDevices]=React.useState([])
    React.useEffect(()=>{
        getPages().then((res)=>{
            setDevices(res)
        })
    },[])
    return (<>
    <div>{'{{.name}}'}</div>
    {devices}
    </>)
}

export default IndexPage
