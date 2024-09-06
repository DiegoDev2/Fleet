package main

// Libxcomposite - X.Org: Client library for the Composite extension
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibxcomposite() {
	// Método 1: Descargar y extraer .tar.gz
	libxcomposite_tar_url := "https://www.x.org/archive/individual/lib/libXcomposite-0.4.6.tar.xz"
	libxcomposite_cmd_tar := exec.Command("curl", "-L", libxcomposite_tar_url, "-o", "package.tar.gz")
	err := libxcomposite_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxcomposite_zip_url := "https://www.x.org/archive/individual/lib/libXcomposite-0.4.6.tar.xz"
	libxcomposite_cmd_zip := exec.Command("curl", "-L", libxcomposite_zip_url, "-o", "package.zip")
	err = libxcomposite_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxcomposite_bin_url := "https://www.x.org/archive/individual/lib/libXcomposite-0.4.6.tar.xz"
	libxcomposite_cmd_bin := exec.Command("curl", "-L", libxcomposite_bin_url, "-o", "binary.bin")
	err = libxcomposite_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxcomposite_src_url := "https://www.x.org/archive/individual/lib/libXcomposite-0.4.6.tar.xz"
	libxcomposite_cmd_src := exec.Command("curl", "-L", libxcomposite_src_url, "-o", "source.tar.gz")
	err = libxcomposite_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxcomposite_cmd_direct := exec.Command("./binary")
	err = libxcomposite_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libxfixes")
exec.Command("latte", "install", "libxfixes")
	fmt.Println("Instalando dependencia: xorgproto")
exec.Command("latte", "install", "xorgproto")
}
