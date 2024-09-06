package main

// Cdktf - Cloud Development Kit for Terraform
// Homepage: https://github.com/hashicorp/terraform-cdk

import (
	"fmt"
	
	"os/exec"
)

func installCdktf() {
	// Método 1: Descargar y extraer .tar.gz
	cdktf_tar_url := "https://registry.npmjs.org/cdktf-cli/-/cdktf-cli-0.20.4.tgz"
	cdktf_cmd_tar := exec.Command("curl", "-L", cdktf_tar_url, "-o", "package.tar.gz")
	err := cdktf_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cdktf_zip_url := "https://registry.npmjs.org/cdktf-cli/-/cdktf-cli-0.20.4.tgz"
	cdktf_cmd_zip := exec.Command("curl", "-L", cdktf_zip_url, "-o", "package.zip")
	err = cdktf_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cdktf_bin_url := "https://registry.npmjs.org/cdktf-cli/-/cdktf-cli-0.20.4.tgz"
	cdktf_cmd_bin := exec.Command("curl", "-L", cdktf_bin_url, "-o", "binary.bin")
	err = cdktf_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cdktf_src_url := "https://registry.npmjs.org/cdktf-cli/-/cdktf-cli-0.20.4.tgz"
	cdktf_cmd_src := exec.Command("curl", "-L", cdktf_src_url, "-o", "source.tar.gz")
	err = cdktf_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cdktf_cmd_direct := exec.Command("./binary")
	err = cdktf_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: node")
exec.Command("latte", "install", "node")
	fmt.Println("Instalando dependencia: terraform")
exec.Command("latte", "install", "terraform")
}
