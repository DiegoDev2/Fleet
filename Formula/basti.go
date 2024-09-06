package main

// Basti - Securely connect to RDS, Elasticache, and other AWS resources in VPCs
// Homepage: https://github.com/basti-app/basti

import (
	"fmt"
	
	"os/exec"
)

func installBasti() {
	// Método 1: Descargar y extraer .tar.gz
	basti_tar_url := "https://registry.npmjs.org/basti/-/basti-1.6.2.tgz"
	basti_cmd_tar := exec.Command("curl", "-L", basti_tar_url, "-o", "package.tar.gz")
	err := basti_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	basti_zip_url := "https://registry.npmjs.org/basti/-/basti-1.6.2.tgz"
	basti_cmd_zip := exec.Command("curl", "-L", basti_zip_url, "-o", "package.zip")
	err = basti_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	basti_bin_url := "https://registry.npmjs.org/basti/-/basti-1.6.2.tgz"
	basti_cmd_bin := exec.Command("curl", "-L", basti_bin_url, "-o", "binary.bin")
	err = basti_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	basti_src_url := "https://registry.npmjs.org/basti/-/basti-1.6.2.tgz"
	basti_cmd_src := exec.Command("curl", "-L", basti_src_url, "-o", "source.tar.gz")
	err = basti_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	basti_cmd_direct := exec.Command("./binary")
	err = basti_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
}
