package main

// Fastnetmon - DDoS detection tool with sFlow, Netflow, IPFIX and port mirror support
// Homepage: https://github.com/pavel-odintsov/fastnetmon/

import (
	"fmt"
	
	"os/exec"
)

func installFastnetmon() {
	// Método 1: Descargar y extraer .tar.gz
	fastnetmon_tar_url := "https://github.com/pavel-odintsov/fastnetmon/archive/refs/tags/v1.2.6.tar.gz"
	fastnetmon_cmd_tar := exec.Command("curl", "-L", fastnetmon_tar_url, "-o", "package.tar.gz")
	err := fastnetmon_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fastnetmon_zip_url := "https://github.com/pavel-odintsov/fastnetmon/archive/refs/tags/v1.2.6.zip"
	fastnetmon_cmd_zip := exec.Command("curl", "-L", fastnetmon_zip_url, "-o", "package.zip")
	err = fastnetmon_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fastnetmon_bin_url := "https://github.com/pavel-odintsov/fastnetmon/archive/refs/tags/v1.2.6.bin"
	fastnetmon_cmd_bin := exec.Command("curl", "-L", fastnetmon_bin_url, "-o", "binary.bin")
	err = fastnetmon_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fastnetmon_src_url := "https://github.com/pavel-odintsov/fastnetmon/archive/refs/tags/v1.2.6.src.tar.gz"
	fastnetmon_cmd_src := exec.Command("curl", "-L", fastnetmon_src_url, "-o", "source.tar.gz")
	err = fastnetmon_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fastnetmon_cmd_direct := exec.Command("./binary")
	err = fastnetmon_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: abseil")
	exec.Command("latte", "install", "abseil").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: capnp")
	exec.Command("latte", "install", "capnp").Run()
	fmt.Println("Instalando dependencia: grpc")
	exec.Command("latte", "install", "grpc").Run()
	fmt.Println("Instalando dependencia: hiredis")
	exec.Command("latte", "install", "hiredis").Run()
	fmt.Println("Instalando dependencia: log4cpp")
	exec.Command("latte", "install", "log4cpp").Run()
	fmt.Println("Instalando dependencia: mongo-c-driver")
	exec.Command("latte", "install", "mongo-c-driver").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: protobuf")
	exec.Command("latte", "install", "protobuf").Run()
	fmt.Println("Instalando dependencia: elfutils")
	exec.Command("latte", "install", "elfutils").Run()
	fmt.Println("Instalando dependencia: libbpf")
	exec.Command("latte", "install", "libbpf").Run()
	fmt.Println("Instalando dependencia: libpcap")
	exec.Command("latte", "install", "libpcap").Run()
}
