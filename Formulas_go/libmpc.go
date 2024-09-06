package main

// Libmpc - C library for the arithmetic of high precision complex numbers
// Homepage: https://www.multiprecision.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibmpc() {
	// Método 1: Descargar y extraer .tar.gz
	libmpc_tar_url := "https://ftp.gnu.org/gnu/mpc/mpc-1.3.1.tar.gz"
	libmpc_cmd_tar := exec.Command("curl", "-L", libmpc_tar_url, "-o", "package.tar.gz")
	err := libmpc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmpc_zip_url := "https://ftp.gnu.org/gnu/mpc/mpc-1.3.1.zip"
	libmpc_cmd_zip := exec.Command("curl", "-L", libmpc_zip_url, "-o", "package.zip")
	err = libmpc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmpc_bin_url := "https://ftp.gnu.org/gnu/mpc/mpc-1.3.1.bin"
	libmpc_cmd_bin := exec.Command("curl", "-L", libmpc_bin_url, "-o", "binary.bin")
	err = libmpc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmpc_src_url := "https://ftp.gnu.org/gnu/mpc/mpc-1.3.1.src.tar.gz"
	libmpc_cmd_src := exec.Command("curl", "-L", libmpc_src_url, "-o", "source.tar.gz")
	err = libmpc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmpc_cmd_direct := exec.Command("./binary")
	err = libmpc_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: mpfr")
exec.Command("latte", "install", "mpfr")
}
