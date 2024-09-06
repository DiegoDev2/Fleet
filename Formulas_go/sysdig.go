package main

// Sysdig - System-level exploration and troubleshooting tool
// Homepage: https://sysdig.com/

import (
	"fmt"
	
	"os/exec"
)

func installSysdig() {
	// Método 1: Descargar y extraer .tar.gz
	sysdig_tar_url := "https://github.com/draios/sysdig/archive/refs/tags/0.38.1.tar.gz"
	sysdig_cmd_tar := exec.Command("curl", "-L", sysdig_tar_url, "-o", "package.tar.gz")
	err := sysdig_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	sysdig_zip_url := "https://github.com/draios/sysdig/archive/refs/tags/0.38.1.zip"
	sysdig_cmd_zip := exec.Command("curl", "-L", sysdig_zip_url, "-o", "package.zip")
	err = sysdig_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	sysdig_bin_url := "https://github.com/draios/sysdig/archive/refs/tags/0.38.1.bin"
	sysdig_cmd_bin := exec.Command("curl", "-L", sysdig_bin_url, "-o", "binary.bin")
	err = sysdig_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	sysdig_src_url := "https://github.com/draios/sysdig/archive/refs/tags/0.38.1.src.tar.gz"
	sysdig_cmd_src := exec.Command("curl", "-L", sysdig_src_url, "-o", "source.tar.gz")
	err = sysdig_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	sysdig_cmd_direct := exec.Command("./binary")
	err = sysdig_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: nlohmann-json")
exec.Command("latte", "install", "nlohmann-json")
	fmt.Println("Instalando dependencia: valijson")
exec.Command("latte", "install", "valijson")
	fmt.Println("Instalando dependencia: jsoncpp")
exec.Command("latte", "install", "jsoncpp")
	fmt.Println("Instalando dependencia: luajit")
exec.Command("latte", "install", "luajit")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
	fmt.Println("Instalando dependencia: re2")
exec.Command("latte", "install", "re2")
	fmt.Println("Instalando dependencia: tbb")
exec.Command("latte", "install", "tbb")
	fmt.Println("Instalando dependencia: uthash")
exec.Command("latte", "install", "uthash")
	fmt.Println("Instalando dependencia: yaml-cpp")
exec.Command("latte", "install", "yaml-cpp")
	fmt.Println("Instalando dependencia: libb64")
exec.Command("latte", "install", "libb64")
	fmt.Println("Instalando dependencia: abseil")
exec.Command("latte", "install", "abseil")
	fmt.Println("Instalando dependencia: elfutils")
exec.Command("latte", "install", "elfutils")
	fmt.Println("Instalando dependencia: grpc")
exec.Command("latte", "install", "grpc")
	fmt.Println("Instalando dependencia: jq")
exec.Command("latte", "install", "jq")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: protobuf")
exec.Command("latte", "install", "protobuf")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
}
