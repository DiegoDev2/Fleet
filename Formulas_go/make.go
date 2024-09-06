package main

// Make - Utility for directing compilation
// Homepage: https://www.gnu.org/software/make/

import (
	"fmt"
	
	"os/exec"
)

func installMake() {
	// Método 1: Descargar y extraer .tar.gz
	make_tar_url := "https://ftp.gnu.org/gnu/make/make-4.4.1.tar.lz"
	make_cmd_tar := exec.Command("curl", "-L", make_tar_url, "-o", "package.tar.gz")
	err := make_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	make_zip_url := "https://ftp.gnu.org/gnu/make/make-4.4.1.tar.lz"
	make_cmd_zip := exec.Command("curl", "-L", make_zip_url, "-o", "package.zip")
	err = make_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	make_bin_url := "https://ftp.gnu.org/gnu/make/make-4.4.1.tar.lz"
	make_cmd_bin := exec.Command("curl", "-L", make_bin_url, "-o", "binary.bin")
	err = make_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	make_src_url := "https://ftp.gnu.org/gnu/make/make-4.4.1.tar.lz"
	make_cmd_src := exec.Command("curl", "-L", make_src_url, "-o", "source.tar.gz")
	err = make_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	make_cmd_direct := exec.Command("./binary")
	err = make_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: texinfo")
exec.Command("latte", "install", "texinfo")
	fmt.Println("Instalando dependencia: wget")
exec.Command("latte", "install", "wget")
}
