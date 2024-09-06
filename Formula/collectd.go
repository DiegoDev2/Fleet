package main

// Collectd - Statistics collection and monitoring daemon
// Homepage: https://collectd.org/

import (
	"fmt"
	
	"os/exec"
)

func installCollectd() {
	// Método 1: Descargar y extraer .tar.gz
	collectd_tar_url := "https://storage.googleapis.com/collectd-tarballs/collectd-5.12.0.tar.bz2"
	collectd_cmd_tar := exec.Command("curl", "-L", collectd_tar_url, "-o", "package.tar.gz")
	err := collectd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	collectd_zip_url := "https://storage.googleapis.com/collectd-tarballs/collectd-5.12.0.tar.bz2"
	collectd_cmd_zip := exec.Command("curl", "-L", collectd_zip_url, "-o", "package.zip")
	err = collectd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	collectd_bin_url := "https://storage.googleapis.com/collectd-tarballs/collectd-5.12.0.tar.bz2"
	collectd_cmd_bin := exec.Command("curl", "-L", collectd_bin_url, "-o", "binary.bin")
	err = collectd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	collectd_src_url := "https://storage.googleapis.com/collectd-tarballs/collectd-5.12.0.tar.bz2"
	collectd_cmd_src := exec.Command("curl", "-L", collectd_src_url, "-o", "source.tar.gz")
	err = collectd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	collectd_cmd_direct := exec.Command("./binary")
	err = collectd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libgcrypt")
	exec.Command("latte", "install", "libgcrypt").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: net-snmp")
	exec.Command("latte", "install", "net-snmp").Run()
	fmt.Println("Instalando dependencia: protobuf-c")
	exec.Command("latte", "install", "protobuf-c").Run()
	fmt.Println("Instalando dependencia: riemann-client")
	exec.Command("latte", "install", "riemann-client").Run()
	fmt.Println("Instalando dependencia: libgpg-error")
	exec.Command("latte", "install", "libgpg-error").Run()
}
