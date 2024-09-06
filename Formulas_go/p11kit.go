package main

// P11Kit - Library to load and enumerate PKCS#11 modules
// Homepage: https://p11-glue.freedesktop.org

import (
	"fmt"
	
	"os/exec"
)

func installP11Kit() {
	// Método 1: Descargar y extraer .tar.gz
	p11kit_tar_url := "https://github.com/p11-glue/p11-kit/releases/download/0.25.5/p11-kit-0.25.5.tar.xz"
	p11kit_cmd_tar := exec.Command("curl", "-L", p11kit_tar_url, "-o", "package.tar.gz")
	err := p11kit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	p11kit_zip_url := "https://github.com/p11-glue/p11-kit/releases/download/0.25.5/p11-kit-0.25.5.tar.xz"
	p11kit_cmd_zip := exec.Command("curl", "-L", p11kit_zip_url, "-o", "package.zip")
	err = p11kit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	p11kit_bin_url := "https://github.com/p11-glue/p11-kit/releases/download/0.25.5/p11-kit-0.25.5.tar.xz"
	p11kit_cmd_bin := exec.Command("curl", "-L", p11kit_bin_url, "-o", "binary.bin")
	err = p11kit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	p11kit_src_url := "https://github.com/p11-glue/p11-kit/releases/download/0.25.5/p11-kit-0.25.5.tar.xz"
	p11kit_cmd_src := exec.Command("curl", "-L", p11kit_src_url, "-o", "source.tar.gz")
	err = p11kit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	p11kit_cmd_direct := exec.Command("./binary")
	err = p11kit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: ca-certificates")
exec.Command("latte", "install", "ca-certificates")
	fmt.Println("Instalando dependencia: libtasn1")
exec.Command("latte", "install", "libtasn1")
}
