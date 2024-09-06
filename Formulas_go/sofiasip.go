package main

// SofiaSip - SIP User-Agent library
// Homepage: https://sofia-sip.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installSofiaSip() {
	// Método 1: Descargar y extraer .tar.gz
	sofiasip_tar_url := "https://github.com/freeswitch/sofia-sip/archive/refs/tags/v1.13.17.tar.gz"
	sofiasip_cmd_tar := exec.Command("curl", "-L", sofiasip_tar_url, "-o", "package.tar.gz")
	err := sofiasip_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sofiasip_zip_url := "https://github.com/freeswitch/sofia-sip/archive/refs/tags/v1.13.17.zip"
	sofiasip_cmd_zip := exec.Command("curl", "-L", sofiasip_zip_url, "-o", "package.zip")
	err = sofiasip_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sofiasip_bin_url := "https://github.com/freeswitch/sofia-sip/archive/refs/tags/v1.13.17.bin"
	sofiasip_cmd_bin := exec.Command("curl", "-L", sofiasip_bin_url, "-o", "binary.bin")
	err = sofiasip_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sofiasip_src_url := "https://github.com/freeswitch/sofia-sip/archive/refs/tags/v1.13.17.src.tar.gz"
	sofiasip_cmd_src := exec.Command("curl", "-L", sofiasip_src_url, "-o", "source.tar.gz")
	err = sofiasip_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sofiasip_cmd_direct := exec.Command("./binary")
	err = sofiasip_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}
