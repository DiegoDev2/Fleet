package main

// Tvnamer - Automatic TV episode file renamer that uses data from thetvdb.com
// Homepage: https://github.com/dbr/tvnamer

import (
	"fmt"
	
	"os/exec"
)

func installTvnamer() {
	// Método 1: Descargar y extraer .tar.gz
	tvnamer_tar_url := "https://files.pythonhosted.org/packages/7e/07/688dc96a86cf212ffdb291d2f012bc4a41ee78324a2eda4c98f05f5e3062/tvnamer-3.0.4.tar.gz"
	tvnamer_cmd_tar := exec.Command("curl", "-L", tvnamer_tar_url, "-o", "package.tar.gz")
	err := tvnamer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tvnamer_zip_url := "https://files.pythonhosted.org/packages/7e/07/688dc96a86cf212ffdb291d2f012bc4a41ee78324a2eda4c98f05f5e3062/tvnamer-3.0.4.zip"
	tvnamer_cmd_zip := exec.Command("curl", "-L", tvnamer_zip_url, "-o", "package.zip")
	err = tvnamer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tvnamer_bin_url := "https://files.pythonhosted.org/packages/7e/07/688dc96a86cf212ffdb291d2f012bc4a41ee78324a2eda4c98f05f5e3062/tvnamer-3.0.4.bin"
	tvnamer_cmd_bin := exec.Command("curl", "-L", tvnamer_bin_url, "-o", "binary.bin")
	err = tvnamer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tvnamer_src_url := "https://files.pythonhosted.org/packages/7e/07/688dc96a86cf212ffdb291d2f012bc4a41ee78324a2eda4c98f05f5e3062/tvnamer-3.0.4.src.tar.gz"
	tvnamer_cmd_src := exec.Command("curl", "-L", tvnamer_src_url, "-o", "source.tar.gz")
	err = tvnamer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tvnamer_cmd_direct := exec.Command("./binary")
	err = tvnamer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
