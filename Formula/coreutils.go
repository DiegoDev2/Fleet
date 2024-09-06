package main

// Coreutils - GNU File, Shell, and Text utilities
// Homepage: https://www.gnu.org/software/coreutils/

import (
	"fmt"
	
	"os/exec"
)

func installCoreutils() {
	// Método 1: Descargar y extraer .tar.gz
	coreutils_tar_url := "https://ftp.gnu.org/gnu/coreutils/coreutils-9.5.tar.xz"
	coreutils_cmd_tar := exec.Command("curl", "-L", coreutils_tar_url, "-o", "package.tar.gz")
	err := coreutils_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	coreutils_zip_url := "https://ftp.gnu.org/gnu/coreutils/coreutils-9.5.tar.xz"
	coreutils_cmd_zip := exec.Command("curl", "-L", coreutils_zip_url, "-o", "package.zip")
	err = coreutils_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	coreutils_bin_url := "https://ftp.gnu.org/gnu/coreutils/coreutils-9.5.tar.xz"
	coreutils_cmd_bin := exec.Command("curl", "-L", coreutils_bin_url, "-o", "binary.bin")
	err = coreutils_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	coreutils_src_url := "https://ftp.gnu.org/gnu/coreutils/coreutils-9.5.tar.xz"
	coreutils_cmd_src := exec.Command("curl", "-L", coreutils_src_url, "-o", "source.tar.gz")
	err = coreutils_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	coreutils_cmd_direct := exec.Command("./binary")
	err = coreutils_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: bison")
	exec.Command("latte", "install", "bison").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: texinfo")
	exec.Command("latte", "install", "texinfo").Run()
	fmt.Println("Instalando dependencia: wget")
	exec.Command("latte", "install", "wget").Run()
	fmt.Println("Instalando dependencia: xz")
	exec.Command("latte", "install", "xz").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: attr")
	exec.Command("latte", "install", "attr").Run()
}
