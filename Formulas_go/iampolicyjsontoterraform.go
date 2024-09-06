package main

// IamPolicyJsonToTerraform - Convert a JSON IAM Policy into terraform
// Homepage: https://github.com/flosell/iam-policy-json-to-terraform

import (
	"fmt"
	
	"os/exec"
)

func installIamPolicyJsonToTerraform() {
	// Método 1: Descargar y extraer .tar.gz
	iampolicyjsontoterraform_tar_url := "https://github.com/flosell/iam-policy-json-to-terraform/archive/refs/tags/1.8.2.tar.gz"
	iampolicyjsontoterraform_cmd_tar := exec.Command("curl", "-L", iampolicyjsontoterraform_tar_url, "-o", "package.tar.gz")
	err := iampolicyjsontoterraform_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	iampolicyjsontoterraform_zip_url := "https://github.com/flosell/iam-policy-json-to-terraform/archive/refs/tags/1.8.2.zip"
	iampolicyjsontoterraform_cmd_zip := exec.Command("curl", "-L", iampolicyjsontoterraform_zip_url, "-o", "package.zip")
	err = iampolicyjsontoterraform_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	iampolicyjsontoterraform_bin_url := "https://github.com/flosell/iam-policy-json-to-terraform/archive/refs/tags/1.8.2.bin"
	iampolicyjsontoterraform_cmd_bin := exec.Command("curl", "-L", iampolicyjsontoterraform_bin_url, "-o", "binary.bin")
	err = iampolicyjsontoterraform_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	iampolicyjsontoterraform_src_url := "https://github.com/flosell/iam-policy-json-to-terraform/archive/refs/tags/1.8.2.src.tar.gz"
	iampolicyjsontoterraform_cmd_src := exec.Command("curl", "-L", iampolicyjsontoterraform_src_url, "-o", "source.tar.gz")
	err = iampolicyjsontoterraform_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	iampolicyjsontoterraform_cmd_direct := exec.Command("./binary")
	err = iampolicyjsontoterraform_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
}
