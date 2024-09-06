package main

// Toot - Mastodon CLI & TUI
// Homepage: https://toot.bezdomni.net/

import (
	"fmt"
	
	"os/exec"
)

func installToot() {
	// Método 1: Descargar y extraer .tar.gz
	toot_tar_url := "https://files.pythonhosted.org/packages/81/be/185a52bc194f00e02875b3dd3161ac8022e5cd6c365b5e90baa65c5ce229/toot-0.44.1.tar.gz"
	toot_cmd_tar := exec.Command("curl", "-L", toot_tar_url, "-o", "package.tar.gz")
	err := toot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	toot_zip_url := "https://files.pythonhosted.org/packages/81/be/185a52bc194f00e02875b3dd3161ac8022e5cd6c365b5e90baa65c5ce229/toot-0.44.1.zip"
	toot_cmd_zip := exec.Command("curl", "-L", toot_zip_url, "-o", "package.zip")
	err = toot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	toot_bin_url := "https://files.pythonhosted.org/packages/81/be/185a52bc194f00e02875b3dd3161ac8022e5cd6c365b5e90baa65c5ce229/toot-0.44.1.bin"
	toot_cmd_bin := exec.Command("curl", "-L", toot_bin_url, "-o", "binary.bin")
	err = toot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	toot_src_url := "https://files.pythonhosted.org/packages/81/be/185a52bc194f00e02875b3dd3161ac8022e5cd6c365b5e90baa65c5ce229/toot-0.44.1.src.tar.gz"
	toot_cmd_src := exec.Command("curl", "-L", toot_src_url, "-o", "source.tar.gz")
	err = toot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	toot_cmd_direct := exec.Command("./binary")
	err = toot_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: certifi")
exec.Command("latte", "install", "certifi")
	fmt.Println("Instalando dependencia: pillow")
exec.Command("latte", "install", "pillow")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
