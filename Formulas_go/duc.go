package main

// Duc - Suite of tools for inspecting disk usage
// Homepage: https://github.com/zevv/duc

import (
	"fmt"
	
	"os/exec"
)

func installDuc() {
	// Método 1: Descargar y extraer .tar.gz
	duc_tar_url := "https://github.com/zevv/duc/releases/download/1.4.5/duc-1.4.5.tar.gz"
	duc_cmd_tar := exec.Command("curl", "-L", duc_tar_url, "-o", "package.tar.gz")
	err := duc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	duc_zip_url := "https://github.com/zevv/duc/releases/download/1.4.5/duc-1.4.5.zip"
	duc_cmd_zip := exec.Command("curl", "-L", duc_zip_url, "-o", "package.zip")
	err = duc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	duc_bin_url := "https://github.com/zevv/duc/releases/download/1.4.5/duc-1.4.5.bin"
	duc_cmd_bin := exec.Command("curl", "-L", duc_bin_url, "-o", "binary.bin")
	err = duc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	duc_src_url := "https://github.com/zevv/duc/releases/download/1.4.5/duc-1.4.5.src.tar.gz"
	duc_cmd_src := exec.Command("curl", "-L", duc_src_url, "-o", "source.tar.gz")
	err = duc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	duc_cmd_direct := exec.Command("./binary")
	err = duc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: glfw")
exec.Command("latte", "install", "glfw")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
	fmt.Println("Instalando dependencia: tokyo-cabinet")
exec.Command("latte", "install", "tokyo-cabinet")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
}
