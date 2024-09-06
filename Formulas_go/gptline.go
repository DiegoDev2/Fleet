package main

// Gptline - ChatGPT client with native iTerm2 support
// Homepage: https://github.com/gnachman/gptline

import (
	"fmt"
	
	"os/exec"
)

func installGptline() {
	// Método 1: Descargar y extraer .tar.gz
	gptline_tar_url := "https://files.pythonhosted.org/packages/5b/28/d15a9a5b349c77a051a633e13141151314f352067ec7d516220bd6b20fcf/gptline-1.0.8.tar.gz"
	gptline_cmd_tar := exec.Command("curl", "-L", gptline_tar_url, "-o", "package.tar.gz")
	err := gptline_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gptline_zip_url := "https://files.pythonhosted.org/packages/5b/28/d15a9a5b349c77a051a633e13141151314f352067ec7d516220bd6b20fcf/gptline-1.0.8.zip"
	gptline_cmd_zip := exec.Command("curl", "-L", gptline_zip_url, "-o", "package.zip")
	err = gptline_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gptline_bin_url := "https://files.pythonhosted.org/packages/5b/28/d15a9a5b349c77a051a633e13141151314f352067ec7d516220bd6b20fcf/gptline-1.0.8.bin"
	gptline_cmd_bin := exec.Command("curl", "-L", gptline_bin_url, "-o", "binary.bin")
	err = gptline_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gptline_src_url := "https://files.pythonhosted.org/packages/5b/28/d15a9a5b349c77a051a633e13141151314f352067ec7d516220bd6b20fcf/gptline-1.0.8.src.tar.gz"
	gptline_cmd_src := exec.Command("curl", "-L", gptline_src_url, "-o", "source.tar.gz")
	err = gptline_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gptline_cmd_direct := exec.Command("./binary")
	err = gptline_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: pillow")
exec.Command("latte", "install", "pillow")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
