package main

// Cloudiscovery - Help you discover resources in the cloud environment
// Homepage: https://github.com/Cloud-Architects/cloudiscovery

import (
	"fmt"
	
	"os/exec"
)

func installCloudiscovery() {
	// Método 1: Descargar y extraer .tar.gz
	cloudiscovery_tar_url := "https://files.pythonhosted.org/packages/d3/c2/9a5f93ac5376f83903c8550bde45e2888da3fb092b63e02e19d6c852134c/cloudiscovery-2.4.4.tar.gz"
	cloudiscovery_cmd_tar := exec.Command("curl", "-L", cloudiscovery_tar_url, "-o", "package.tar.gz")
	err := cloudiscovery_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cloudiscovery_zip_url := "https://files.pythonhosted.org/packages/d3/c2/9a5f93ac5376f83903c8550bde45e2888da3fb092b63e02e19d6c852134c/cloudiscovery-2.4.4.zip"
	cloudiscovery_cmd_zip := exec.Command("curl", "-L", cloudiscovery_zip_url, "-o", "package.zip")
	err = cloudiscovery_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cloudiscovery_bin_url := "https://files.pythonhosted.org/packages/d3/c2/9a5f93ac5376f83903c8550bde45e2888da3fb092b63e02e19d6c852134c/cloudiscovery-2.4.4.bin"
	cloudiscovery_cmd_bin := exec.Command("curl", "-L", cloudiscovery_bin_url, "-o", "binary.bin")
	err = cloudiscovery_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cloudiscovery_src_url := "https://files.pythonhosted.org/packages/d3/c2/9a5f93ac5376f83903c8550bde45e2888da3fb092b63e02e19d6c852134c/cloudiscovery-2.4.4.src.tar.gz"
	cloudiscovery_cmd_src := exec.Command("curl", "-L", cloudiscovery_src_url, "-o", "source.tar.gz")
	err = cloudiscovery_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cloudiscovery_cmd_direct := exec.Command("./binary")
	err = cloudiscovery_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
