package main

// Gcalcli - Easily access your Google Calendar(s) from a command-line
// Homepage: https://github.com/insanum/gcalcli

import (
	"fmt"
	
	"os/exec"
)

func installGcalcli() {
	// Método 1: Descargar y extraer .tar.gz
	gcalcli_tar_url := "https://files.pythonhosted.org/packages/f0/a3/713e440c41a9dcbace59b864097c439c0b7248c5d7bfad8a8b0fcc7ed096/gcalcli-4.4.0.tar.gz"
	gcalcli_cmd_tar := exec.Command("curl", "-L", gcalcli_tar_url, "-o", "package.tar.gz")
	err := gcalcli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gcalcli_zip_url := "https://files.pythonhosted.org/packages/f0/a3/713e440c41a9dcbace59b864097c439c0b7248c5d7bfad8a8b0fcc7ed096/gcalcli-4.4.0.zip"
	gcalcli_cmd_zip := exec.Command("curl", "-L", gcalcli_zip_url, "-o", "package.zip")
	err = gcalcli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gcalcli_bin_url := "https://files.pythonhosted.org/packages/f0/a3/713e440c41a9dcbace59b864097c439c0b7248c5d7bfad8a8b0fcc7ed096/gcalcli-4.4.0.bin"
	gcalcli_cmd_bin := exec.Command("curl", "-L", gcalcli_bin_url, "-o", "binary.bin")
	err = gcalcli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gcalcli_src_url := "https://files.pythonhosted.org/packages/f0/a3/713e440c41a9dcbace59b864097c439c0b7248c5d7bfad8a8b0fcc7ed096/gcalcli-4.4.0.src.tar.gz"
	gcalcli_cmd_src := exec.Command("curl", "-L", gcalcli_src_url, "-o", "source.tar.gz")
	err = gcalcli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gcalcli_cmd_direct := exec.Command("./binary")
	err = gcalcli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
