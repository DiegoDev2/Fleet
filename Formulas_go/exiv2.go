package main

// Exiv2 - EXIF and IPTC metadata manipulation library and tools
// Homepage: https://exiv2.org/

import (
	"fmt"
	
	"os/exec"
)

func installExiv2() {
	// Método 1: Descargar y extraer .tar.gz
	exiv2_tar_url := "https://github.com/Exiv2/exiv2/archive/refs/tags/v0.28.3.tar.gz"
	exiv2_cmd_tar := exec.Command("curl", "-L", exiv2_tar_url, "-o", "package.tar.gz")
	err := exiv2_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	exiv2_zip_url := "https://github.com/Exiv2/exiv2/archive/refs/tags/v0.28.3.zip"
	exiv2_cmd_zip := exec.Command("curl", "-L", exiv2_zip_url, "-o", "package.zip")
	err = exiv2_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	exiv2_bin_url := "https://github.com/Exiv2/exiv2/archive/refs/tags/v0.28.3.bin"
	exiv2_cmd_bin := exec.Command("curl", "-L", exiv2_bin_url, "-o", "binary.bin")
	err = exiv2_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	exiv2_src_url := "https://github.com/Exiv2/exiv2/archive/refs/tags/v0.28.3.src.tar.gz"
	exiv2_cmd_src := exec.Command("curl", "-L", exiv2_src_url, "-o", "source.tar.gz")
	err = exiv2_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	exiv2_cmd_direct := exec.Command("./binary")
	err = exiv2_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: brotli")
exec.Command("latte", "install", "brotli")
	fmt.Println("Instalando dependencia: inih")
exec.Command("latte", "install", "inih")
	fmt.Println("Instalando dependencia: libssh")
exec.Command("latte", "install", "libssh")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
