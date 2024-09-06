package main

// OpentelemetryCpp - OpenTelemetry C++ Client
// Homepage: https://opentelemetry.io/

import (
	"fmt"
	
	"os/exec"
)

func installOpentelemetryCpp() {
	// Método 1: Descargar y extraer .tar.gz
	opentelemetrycpp_tar_url := "https://github.com/open-telemetry/opentelemetry-cpp/archive/refs/tags/v1.16.1.tar.gz"
	opentelemetrycpp_cmd_tar := exec.Command("curl", "-L", opentelemetrycpp_tar_url, "-o", "package.tar.gz")
	err := opentelemetrycpp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	opentelemetrycpp_zip_url := "https://github.com/open-telemetry/opentelemetry-cpp/archive/refs/tags/v1.16.1.zip"
	opentelemetrycpp_cmd_zip := exec.Command("curl", "-L", opentelemetrycpp_zip_url, "-o", "package.zip")
	err = opentelemetrycpp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	opentelemetrycpp_bin_url := "https://github.com/open-telemetry/opentelemetry-cpp/archive/refs/tags/v1.16.1.bin"
	opentelemetrycpp_cmd_bin := exec.Command("curl", "-L", opentelemetrycpp_bin_url, "-o", "binary.bin")
	err = opentelemetrycpp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	opentelemetrycpp_src_url := "https://github.com/open-telemetry/opentelemetry-cpp/archive/refs/tags/v1.16.1.src.tar.gz"
	opentelemetrycpp_cmd_src := exec.Command("curl", "-L", opentelemetrycpp_src_url, "-o", "source.tar.gz")
	err = opentelemetrycpp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	opentelemetrycpp_cmd_direct := exec.Command("./binary")
	err = opentelemetrycpp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: abseil")
exec.Command("latte", "install", "abseil")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: grpc")
exec.Command("latte", "install", "grpc")
	fmt.Println("Instalando dependencia: nlohmann-json")
exec.Command("latte", "install", "nlohmann-json")
	fmt.Println("Instalando dependencia: prometheus-cpp")
exec.Command("latte", "install", "prometheus-cpp")
	fmt.Println("Instalando dependencia: protobuf")
exec.Command("latte", "install", "protobuf")
	fmt.Println("Instalando dependencia: c-ares")
exec.Command("latte", "install", "c-ares")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
	fmt.Println("Instalando dependencia: re2")
exec.Command("latte", "install", "re2")
}
