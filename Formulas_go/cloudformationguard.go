package main

// CloudformationGuard - Checks CloudFormation templates for compliance using a declarative syntax
// Homepage: https://github.com/aws-cloudformation/cloudformation-guard

import (
	"fmt"
	
	"os/exec"
)

func installCloudformationGuard() {
	// Método 1: Descargar y extraer .tar.gz
	cloudformationguard_tar_url := "https://github.com/aws-cloudformation/cloudformation-guard/archive/refs/tags/3.1.1.tar.gz"
	cloudformationguard_cmd_tar := exec.Command("curl", "-L", cloudformationguard_tar_url, "-o", "package.tar.gz")
	err := cloudformationguard_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cloudformationguard_zip_url := "https://github.com/aws-cloudformation/cloudformation-guard/archive/refs/tags/3.1.1.zip"
	cloudformationguard_cmd_zip := exec.Command("curl", "-L", cloudformationguard_zip_url, "-o", "package.zip")
	err = cloudformationguard_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cloudformationguard_bin_url := "https://github.com/aws-cloudformation/cloudformation-guard/archive/refs/tags/3.1.1.bin"
	cloudformationguard_cmd_bin := exec.Command("curl", "-L", cloudformationguard_bin_url, "-o", "binary.bin")
	err = cloudformationguard_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cloudformationguard_src_url := "https://github.com/aws-cloudformation/cloudformation-guard/archive/refs/tags/3.1.1.src.tar.gz"
	cloudformationguard_cmd_src := exec.Command("curl", "-L", cloudformationguard_src_url, "-o", "source.tar.gz")
	err = cloudformationguard_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cloudformationguard_cmd_direct := exec.Command("./binary")
	err = cloudformationguard_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
}
