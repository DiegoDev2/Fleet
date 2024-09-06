package main

// Mentat - Coding assistant that leverages GPT-4 to write code
// Homepage: https://www.mentat.ai

import (
	"fmt"
	
	"os/exec"
)

func installMentat() {
	// Método 1: Descargar y extraer .tar.gz
	mentat_tar_url := "https://files.pythonhosted.org/packages/e0/8b/a808d6663065e3b446d3be521d7836f774f3b39bdd30f786d093aca383b6/mentat-1.0.8.tar.gz"
	mentat_cmd_tar := exec.Command("curl", "-L", mentat_tar_url, "-o", "package.tar.gz")
	err := mentat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mentat_zip_url := "https://files.pythonhosted.org/packages/e0/8b/a808d6663065e3b446d3be521d7836f774f3b39bdd30f786d093aca383b6/mentat-1.0.8.zip"
	mentat_cmd_zip := exec.Command("curl", "-L", mentat_zip_url, "-o", "package.zip")
	err = mentat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mentat_bin_url := "https://files.pythonhosted.org/packages/e0/8b/a808d6663065e3b446d3be521d7836f774f3b39bdd30f786d093aca383b6/mentat-1.0.8.bin"
	mentat_cmd_bin := exec.Command("curl", "-L", mentat_bin_url, "-o", "binary.bin")
	err = mentat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mentat_src_url := "https://files.pythonhosted.org/packages/e0/8b/a808d6663065e3b446d3be521d7836f774f3b39bdd30f786d093aca383b6/mentat-1.0.8.src.tar.gz"
	mentat_cmd_src := exec.Command("curl", "-L", mentat_src_url, "-o", "source.tar.gz")
	err = mentat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mentat_cmd_direct := exec.Command("./binary")
	err = mentat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: numpy")
exec.Command("latte", "install", "numpy")
	fmt.Println("Instalando dependencia: pillow")
exec.Command("latte", "install", "pillow")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
