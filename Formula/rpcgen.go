package main

// Rpcgen - Protocol Compiler
// Homepage: https://opensource.apple.com/

import (
	"fmt"
	
	"os/exec"
)

func installRpcgen() {
	// Método 1: Descargar y extraer .tar.gz
	rpcgen_tar_url := "https://github.com/apple-oss-distributions/developer_cmds/archive/refs/tags/developer_cmds-79.tar.gz"
	rpcgen_cmd_tar := exec.Command("curl", "-L", rpcgen_tar_url, "-o", "package.tar.gz")
	err := rpcgen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rpcgen_zip_url := "https://github.com/apple-oss-distributions/developer_cmds/archive/refs/tags/developer_cmds-79.zip"
	rpcgen_cmd_zip := exec.Command("curl", "-L", rpcgen_zip_url, "-o", "package.zip")
	err = rpcgen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rpcgen_bin_url := "https://github.com/apple-oss-distributions/developer_cmds/archive/refs/tags/developer_cmds-79.bin"
	rpcgen_cmd_bin := exec.Command("curl", "-L", rpcgen_bin_url, "-o", "binary.bin")
	err = rpcgen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rpcgen_src_url := "https://github.com/apple-oss-distributions/developer_cmds/archive/refs/tags/developer_cmds-79.src.tar.gz"
	rpcgen_cmd_src := exec.Command("curl", "-L", rpcgen_src_url, "-o", "source.tar.gz")
	err = rpcgen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rpcgen_cmd_direct := exec.Command("./binary")
	err = rpcgen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
