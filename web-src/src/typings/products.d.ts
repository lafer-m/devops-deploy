interface Products {
    name: string
    version?: string[]
}

interface IService {
    name: string
    enabled: boolean
    replicas: number
    value: string
}

interface IPackage {
    type: string
    path: string[]
    selectedPath?: string
}