package main

// Gnunet - Framework for distributed, secure and privacy-preserving applications
// Homepage: https://gnunet.org/

import (
	"fmt"
	
	"os/exec"
)

func installGnunet() {
	// Método 1: Descargar y extraer .tar.gz
	gnunet_tar_url := "https://ftp.gnu.org/gnu/gnunet/gnunet-0.22.0.tar.gz"
	gnunet_cmd_tar := exec.Command("curl", "-L", gnunet_tar_url, "-o", "package.tar.gz")
	err := gnunet_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gnunet_zip_url := "https://ftp.gnu.org/gnu/gnunet/gnunet-0.22.0.zip"
	gnunet_cmd_zip := exec.Command("curl", "-L", gnunet_zip_url, "-o", "package.zip")
	err = gnunet_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gnunet_bin_url := "https://ftp.gnu.org/gnu/gnunet/gnunet-0.22.0.bin"
	gnunet_cmd_bin := exec.Command("curl", "-L", gnunet_bin_url, "-o", "binary.bin")
	err = gnunet_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gnunet_src_url := "https://ftp.gnu.org/gnu/gnunet/gnunet-0.22.0.src.tar.gz"
	gnunet_cmd_src := exec.Command("curl", "-L", gnunet_src_url, "-o", "source.tar.gz")
	err = gnunet_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gnunet_cmd_direct := exec.Command("./binary")
	err = gnunet_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: gnutls")
	exec.Command("latte", "install", "gnutls").Run()
	fmt.Println("Instalando dependencia: jansson")
	exec.Command("latte", "install", "jansson").Run()
	fmt.Println("Instalando dependencia: libextractor")
	exec.Command("latte", "install", "libextractor").Run()
	fmt.Println("Instalando dependencia: libgcrypt")
	exec.Command("latte", "install", "libgcrypt").Run()
	fmt.Println("Instalando dependencia: libidn2")
	exec.Command("latte", "install", "libidn2").Run()
	fmt.Println("Instalando dependencia: libmicrohttpd")
	exec.Command("latte", "install", "libmicrohttpd").Run()
	fmt.Println("Instalando dependencia: libsodium")
	exec.Command("latte", "install", "libsodium").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: libunistring")
	exec.Command("latte", "install", "libunistring").Run()
	fmt.Println("Instalando dependencia: libgpg-error")
	exec.Command("latte", "install", "libgpg-error").Run()
}
