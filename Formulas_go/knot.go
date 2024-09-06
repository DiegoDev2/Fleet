package main

// Knot - High-performance authoritative-only DNS server
// Homepage: https://www.knot-dns.cz/

import (
	"fmt"
	
	"os/exec"
)

func installKnot() {
	// Método 1: Descargar y extraer .tar.gz
	knot_tar_url := "https://secure.nic.cz/files/knot-dns/knot-3.3.5.tar.xz"
	knot_cmd_tar := exec.Command("curl", "-L", knot_tar_url, "-o", "package.tar.gz")
	err := knot_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	knot_zip_url := "https://secure.nic.cz/files/knot-dns/knot-3.3.5.tar.xz"
	knot_cmd_zip := exec.Command("curl", "-L", knot_zip_url, "-o", "package.zip")
	err = knot_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	knot_bin_url := "https://secure.nic.cz/files/knot-dns/knot-3.3.5.tar.xz"
	knot_cmd_bin := exec.Command("curl", "-L", knot_bin_url, "-o", "binary.bin")
	err = knot_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	knot_src_url := "https://secure.nic.cz/files/knot-dns/knot-3.3.5.tar.xz"
	knot_cmd_src := exec.Command("curl", "-L", knot_src_url, "-o", "source.tar.gz")
	err = knot_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	knot_cmd_direct := exec.Command("./binary")
	err = knot_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: sphinx-doc")
exec.Command("latte", "install", "sphinx-doc")
	fmt.Println("Instalando dependencia: fstrm")
exec.Command("latte", "install", "fstrm")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: libidn2")
exec.Command("latte", "install", "libidn2")
	fmt.Println("Instalando dependencia: libnghttp2")
exec.Command("latte", "install", "libnghttp2")
	fmt.Println("Instalando dependencia: lmdb")
exec.Command("latte", "install", "lmdb")
	fmt.Println("Instalando dependencia: protobuf-c")
exec.Command("latte", "install", "protobuf-c")
	fmt.Println("Instalando dependencia: userspace-rcu")
exec.Command("latte", "install", "userspace-rcu")
}
