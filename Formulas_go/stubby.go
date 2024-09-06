package main

// Stubby - DNS privacy enabled stub resolver service based on getdns
// Homepage: https://dnsprivacy.org/wiki/display/DP/DNS+Privacy+Daemon+-+Stubby

import (
	"fmt"
	
	"os/exec"
)

func installStubby() {
	// Método 1: Descargar y extraer .tar.gz
	stubby_tar_url := "https://github.com/getdnsapi/stubby/archive/refs/tags/v0.4.3.tar.gz"
	stubby_cmd_tar := exec.Command("curl", "-L", stubby_tar_url, "-o", "package.tar.gz")
	err := stubby_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stubby_zip_url := "https://github.com/getdnsapi/stubby/archive/refs/tags/v0.4.3.zip"
	stubby_cmd_zip := exec.Command("curl", "-L", stubby_zip_url, "-o", "package.zip")
	err = stubby_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stubby_bin_url := "https://github.com/getdnsapi/stubby/archive/refs/tags/v0.4.3.bin"
	stubby_cmd_bin := exec.Command("curl", "-L", stubby_bin_url, "-o", "binary.bin")
	err = stubby_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stubby_src_url := "https://github.com/getdnsapi/stubby/archive/refs/tags/v0.4.3.src.tar.gz"
	stubby_cmd_src := exec.Command("curl", "-L", stubby_src_url, "-o", "source.tar.gz")
	err = stubby_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stubby_cmd_direct := exec.Command("./binary")
	err = stubby_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: getdns")
exec.Command("latte", "install", "getdns")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: libidn2")
exec.Command("latte", "install", "libidn2")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: unbound")
exec.Command("latte", "install", "unbound")
	fmt.Println("Instalando dependencia: bind")
exec.Command("latte", "install", "bind")
}
