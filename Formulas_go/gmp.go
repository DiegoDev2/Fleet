package main

// Gmp - GNU multiple precision arithmetic library
// Homepage: https://gmplib.org/

import (
	"fmt"
	
	"os/exec"
)

func installGmp() {
	// Método 1: Descargar y extraer .tar.gz
	gmp_tar_url := "https://ftp.gnu.org/gnu/gmp/gmp-6.3.0.tar.xz"
	gmp_cmd_tar := exec.Command("curl", "-L", gmp_tar_url, "-o", "package.tar.gz")
	err := gmp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gmp_zip_url := "https://ftp.gnu.org/gnu/gmp/gmp-6.3.0.tar.xz"
	gmp_cmd_zip := exec.Command("curl", "-L", gmp_zip_url, "-o", "package.zip")
	err = gmp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gmp_bin_url := "https://ftp.gnu.org/gnu/gmp/gmp-6.3.0.tar.xz"
	gmp_cmd_bin := exec.Command("curl", "-L", gmp_bin_url, "-o", "binary.bin")
	err = gmp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gmp_src_url := "https://ftp.gnu.org/gnu/gmp/gmp-6.3.0.tar.xz"
	gmp_cmd_src := exec.Command("curl", "-L", gmp_src_url, "-o", "source.tar.gz")
	err = gmp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gmp_cmd_direct := exec.Command("./binary")
	err = gmp_cmd_direct.Run()
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
}
