package main

// Bzt - BlazeMeter Taurus
// Homepage: https://gettaurus.org/

import (
	"fmt"
	
	"os/exec"
)

func installBzt() {
	// Método 1: Descargar y extraer .tar.gz
	bzt_tar_url := "https://files.pythonhosted.org/packages/94/34/94de1da69ec151e4d0f5c834ad480b46d686d47f1769f4d0f42fb29a636a/bzt-1.16.34.tar.gz"
	bzt_cmd_tar := exec.Command("curl", "-L", bzt_tar_url, "-o", "package.tar.gz")
	err := bzt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bzt_zip_url := "https://files.pythonhosted.org/packages/94/34/94de1da69ec151e4d0f5c834ad480b46d686d47f1769f4d0f42fb29a636a/bzt-1.16.34.zip"
	bzt_cmd_zip := exec.Command("curl", "-L", bzt_zip_url, "-o", "package.zip")
	err = bzt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bzt_bin_url := "https://files.pythonhosted.org/packages/94/34/94de1da69ec151e4d0f5c834ad480b46d686d47f1769f4d0f42fb29a636a/bzt-1.16.34.bin"
	bzt_cmd_bin := exec.Command("curl", "-L", bzt_bin_url, "-o", "binary.bin")
	err = bzt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bzt_src_url := "https://files.pythonhosted.org/packages/94/34/94de1da69ec151e4d0f5c834ad480b46d686d47f1769f4d0f42fb29a636a/bzt-1.16.34.src.tar.gz"
	bzt_cmd_src := exec.Command("curl", "-L", bzt_src_url, "-o", "source.tar.gz")
	err = bzt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bzt_cmd_direct := exec.Command("./binary")
	err = bzt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: certifi")
	exec.Command("latte", "install", "certifi").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: numpy")
	exec.Command("latte", "install", "numpy").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
