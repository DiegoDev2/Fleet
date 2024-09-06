package main

// Silicon - Create beautiful image of your source code
// Homepage: https://github.com/Aloxaf/silicon/

import (
	"fmt"
	
	"os/exec"
)

func installSilicon() {
	// Método 1: Descargar y extraer .tar.gz
	silicon_tar_url := "https://github.com/Aloxaf/silicon/archive/refs/tags/v0.5.2.tar.gz"
	silicon_cmd_tar := exec.Command("curl", "-L", silicon_tar_url, "-o", "package.tar.gz")
	err := silicon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	silicon_zip_url := "https://github.com/Aloxaf/silicon/archive/refs/tags/v0.5.2.zip"
	silicon_cmd_zip := exec.Command("curl", "-L", silicon_zip_url, "-o", "package.zip")
	err = silicon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	silicon_bin_url := "https://github.com/Aloxaf/silicon/archive/refs/tags/v0.5.2.bin"
	silicon_cmd_bin := exec.Command("curl", "-L", silicon_bin_url, "-o", "binary.bin")
	err = silicon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	silicon_src_url := "https://github.com/Aloxaf/silicon/archive/refs/tags/v0.5.2.src.tar.gz"
	silicon_cmd_src := exec.Command("curl", "-L", silicon_src_url, "-o", "source.tar.gz")
	err = silicon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	silicon_cmd_direct := exec.Command("./binary")
	err = silicon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: fontconfig")
exec.Command("latte", "install", "fontconfig")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: libxcb")
exec.Command("latte", "install", "libxcb")
	fmt.Println("Instalando dependencia: xclip")
exec.Command("latte", "install", "xclip")
}
