package main

// Pympress - Simple and powerful dual-screen PDF reader designed for presentations
// Homepage: https://github.com/Cimbali/pympress/

import (
	"fmt"
	
	"os/exec"
)

func installPympress() {
	// Método 1: Descargar y extraer .tar.gz
	pympress_tar_url := "https://files.pythonhosted.org/packages/fb/e2/91827c485aae28d69f0b40c6d366b9f6eb96d8208a98af0345e0ade3fbbd/pympress-1.8.5.tar.gz"
	pympress_cmd_tar := exec.Command("curl", "-L", pympress_tar_url, "-o", "package.tar.gz")
	err := pympress_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	pympress_zip_url := "https://files.pythonhosted.org/packages/fb/e2/91827c485aae28d69f0b40c6d366b9f6eb96d8208a98af0345e0ade3fbbd/pympress-1.8.5.zip"
	pympress_cmd_zip := exec.Command("curl", "-L", pympress_zip_url, "-o", "package.zip")
	err = pympress_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	pympress_bin_url := "https://files.pythonhosted.org/packages/fb/e2/91827c485aae28d69f0b40c6d366b9f6eb96d8208a98af0345e0ade3fbbd/pympress-1.8.5.bin"
	pympress_cmd_bin := exec.Command("curl", "-L", pympress_bin_url, "-o", "binary.bin")
	err = pympress_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	pympress_src_url := "https://files.pythonhosted.org/packages/fb/e2/91827c485aae28d69f0b40c6d366b9f6eb96d8208a98af0345e0ade3fbbd/pympress-1.8.5.src.tar.gz"
	pympress_cmd_src := exec.Command("curl", "-L", pympress_src_url, "-o", "source.tar.gz")
	err = pympress_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	pympress_cmd_direct := exec.Command("./binary")
	err = pympress_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gobject-introspection")
exec.Command("latte", "install", "gobject-introspection")
	fmt.Println("Instalando dependencia: gstreamer")
exec.Command("latte", "install", "gstreamer")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: poppler")
exec.Command("latte", "install", "poppler")
	fmt.Println("Instalando dependencia: pygobject3")
exec.Command("latte", "install", "pygobject3")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
