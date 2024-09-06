package main

// Cubeb - Cross-platform audio library
// Homepage: https://github.com/mozilla/cubeb

import (
	"fmt"
	
	"os/exec"
)

func installCubeb() {
	// Método 1: Descargar y extraer .tar.gz
	cubeb_tar_url := "https://github.com/mozilla/cubeb/archive/refs/tags/cubeb-0.2.tar.gz"
	cubeb_cmd_tar := exec.Command("curl", "-L", cubeb_tar_url, "-o", "package.tar.gz")
	err := cubeb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cubeb_zip_url := "https://github.com/mozilla/cubeb/archive/refs/tags/cubeb-0.2.zip"
	cubeb_cmd_zip := exec.Command("curl", "-L", cubeb_zip_url, "-o", "package.zip")
	err = cubeb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cubeb_bin_url := "https://github.com/mozilla/cubeb/archive/refs/tags/cubeb-0.2.bin"
	cubeb_cmd_bin := exec.Command("curl", "-L", cubeb_bin_url, "-o", "binary.bin")
	err = cubeb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cubeb_src_url := "https://github.com/mozilla/cubeb/archive/refs/tags/cubeb-0.2.src.tar.gz"
	cubeb_cmd_src := exec.Command("curl", "-L", cubeb_src_url, "-o", "source.tar.gz")
	err = cubeb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cubeb_cmd_direct := exec.Command("./binary")
	err = cubeb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: pulseaudio")
	exec.Command("latte", "install", "pulseaudio").Run()
}
