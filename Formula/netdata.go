package main

// Netdata - Diagnose infrastructure problems with metrics, visualizations & alarms
// Homepage: https://netdata.cloud/

import (
	"fmt"
	
	"os/exec"
)

func installNetdata() {
	// Método 1: Descargar y extraer .tar.gz
	netdata_tar_url := "https://github.com/netdata/netdata/releases/download/v1.44.3/netdata-v1.44.3.tar.gz"
	netdata_cmd_tar := exec.Command("curl", "-L", netdata_tar_url, "-o", "package.tar.gz")
	err := netdata_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	netdata_zip_url := "https://github.com/netdata/netdata/releases/download/v1.44.3/netdata-v1.44.3.zip"
	netdata_cmd_zip := exec.Command("curl", "-L", netdata_zip_url, "-o", "package.zip")
	err = netdata_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	netdata_bin_url := "https://github.com/netdata/netdata/releases/download/v1.44.3/netdata-v1.44.3.bin"
	netdata_cmd_bin := exec.Command("curl", "-L", netdata_bin_url, "-o", "binary.bin")
	err = netdata_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	netdata_src_url := "https://github.com/netdata/netdata/releases/download/v1.44.3/netdata-v1.44.3.src.tar.gz"
	netdata_cmd_src := exec.Command("curl", "-L", netdata_src_url, "-o", "source.tar.gz")
	err = netdata_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	netdata_cmd_direct := exec.Command("./binary")
	err = netdata_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: m4")
	exec.Command("latte", "install", "m4").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: abseil")
	exec.Command("latte", "install", "abseil").Run()
	fmt.Println("Instalando dependencia: json-c")
	exec.Command("latte", "install", "json-c").Run()
	fmt.Println("Instalando dependencia: libuv")
	exec.Command("latte", "install", "libuv").Run()
	fmt.Println("Instalando dependencia: libyaml")
	exec.Command("latte", "install", "libyaml").Run()
	fmt.Println("Instalando dependencia: lz4")
	exec.Command("latte", "install", "lz4").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: pcre2")
	exec.Command("latte", "install", "pcre2").Run()
	fmt.Println("Instalando dependencia: protobuf")
	exec.Command("latte", "install", "protobuf").Run()
	fmt.Println("Instalando dependencia: protobuf-c")
	exec.Command("latte", "install", "protobuf-c").Run()
	fmt.Println("Instalando dependencia: util-linux")
	exec.Command("latte", "install", "util-linux").Run()
}
