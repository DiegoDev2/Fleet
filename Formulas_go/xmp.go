package main

// Xmp - Command-line player for module music formats (MOD, S3M, IT, etc)
// Homepage: https://xmp.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installXmp() {
	// Método 1: Descargar y extraer .tar.gz
	xmp_tar_url := "https://github.com/libxmp/xmp-cli/releases/download/xmp-4.2.0/xmp-4.2.0.tar.gz"
	xmp_cmd_tar := exec.Command("curl", "-L", xmp_tar_url, "-o", "package.tar.gz")
	err := xmp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xmp_zip_url := "https://github.com/libxmp/xmp-cli/releases/download/xmp-4.2.0/xmp-4.2.0.zip"
	xmp_cmd_zip := exec.Command("curl", "-L", xmp_zip_url, "-o", "package.zip")
	err = xmp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xmp_bin_url := "https://github.com/libxmp/xmp-cli/releases/download/xmp-4.2.0/xmp-4.2.0.bin"
	xmp_cmd_bin := exec.Command("curl", "-L", xmp_bin_url, "-o", "binary.bin")
	err = xmp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xmp_src_url := "https://github.com/libxmp/xmp-cli/releases/download/xmp-4.2.0/xmp-4.2.0.src.tar.gz"
	xmp_cmd_src := exec.Command("curl", "-L", xmp_src_url, "-o", "source.tar.gz")
	err = xmp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xmp_cmd_direct := exec.Command("./binary")
	err = xmp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libxmp")
exec.Command("latte", "install", "libxmp")
}
