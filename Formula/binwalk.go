package main

// Binwalk - Searches a binary image for embedded files and executable code
// Homepage: https://github.com/ReFirmLabs/binwalk

import (
	"fmt"
	
	"os/exec"
)

func installBinwalk() {
	// Método 1: Descargar y extraer .tar.gz
	binwalk_tar_url := "https://github.com/ReFirmLabs/binwalk/archive/refs/tags/v2.3.4.tar.gz"
	binwalk_cmd_tar := exec.Command("curl", "-L", binwalk_tar_url, "-o", "package.tar.gz")
	err := binwalk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	binwalk_zip_url := "https://github.com/ReFirmLabs/binwalk/archive/refs/tags/v2.3.4.zip"
	binwalk_cmd_zip := exec.Command("curl", "-L", binwalk_zip_url, "-o", "package.zip")
	err = binwalk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	binwalk_bin_url := "https://github.com/ReFirmLabs/binwalk/archive/refs/tags/v2.3.4.bin"
	binwalk_cmd_bin := exec.Command("curl", "-L", binwalk_bin_url, "-o", "binary.bin")
	err = binwalk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	binwalk_src_url := "https://github.com/ReFirmLabs/binwalk/archive/refs/tags/v2.3.4.src.tar.gz"
	binwalk_cmd_src := exec.Command("curl", "-L", binwalk_src_url, "-o", "source.tar.gz")
	err = binwalk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	binwalk_cmd_direct := exec.Command("./binary")
	err = binwalk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: swig")
	exec.Command("latte", "install", "swig").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: numpy")
	exec.Command("latte", "install", "numpy").Run()
	fmt.Println("Instalando dependencia: p7zip")
	exec.Command("latte", "install", "p7zip").Run()
	fmt.Println("Instalando dependencia: pillow")
	exec.Command("latte", "install", "pillow").Run()
	fmt.Println("Instalando dependencia: python@3.11")
	exec.Command("latte", "install", "python@3.11").Run()
	fmt.Println("Instalando dependencia: qhull")
	exec.Command("latte", "install", "qhull").Run()
	fmt.Println("Instalando dependencia: ssdeep")
	exec.Command("latte", "install", "ssdeep").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: patchelf")
	exec.Command("latte", "install", "patchelf").Run()
}
