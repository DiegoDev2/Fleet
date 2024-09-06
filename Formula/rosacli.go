package main

// RosaCli - RedHat OpenShift Service on AWS (ROSA) command-line interface
// Homepage: https://www.openshift.com/products/amazon-openshift

import (
	"fmt"
	
	"os/exec"
)

func installRosaCli() {
	// Método 1: Descargar y extraer .tar.gz
	rosacli_tar_url := "https://github.com/openshift/rosa/archive/refs/tags/v1.2.44.tar.gz"
	rosacli_cmd_tar := exec.Command("curl", "-L", rosacli_tar_url, "-o", "package.tar.gz")
	err := rosacli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rosacli_zip_url := "https://github.com/openshift/rosa/archive/refs/tags/v1.2.44.zip"
	rosacli_cmd_zip := exec.Command("curl", "-L", rosacli_zip_url, "-o", "package.zip")
	err = rosacli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rosacli_bin_url := "https://github.com/openshift/rosa/archive/refs/tags/v1.2.44.bin"
	rosacli_cmd_bin := exec.Command("curl", "-L", rosacli_bin_url, "-o", "binary.bin")
	err = rosacli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rosacli_src_url := "https://github.com/openshift/rosa/archive/refs/tags/v1.2.44.src.tar.gz"
	rosacli_cmd_src := exec.Command("curl", "-L", rosacli_src_url, "-o", "source.tar.gz")
	err = rosacli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rosacli_cmd_direct := exec.Command("./binary")
	err = rosacli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
	exec.Command("latte", "install", "go").Run()
	fmt.Println("Instalando dependencia: awscli")
	exec.Command("latte", "install", "awscli").Run()
}
