package main

// Vineyard - In-memory immutable data manager. (Project under CNCF)
// Homepage: https://v6d.io

import (
	"fmt"
	
	"os/exec"
)

func installVineyard() {
	// Método 1: Descargar y extraer .tar.gz
	vineyard_tar_url := "https://github.com/v6d-io/v6d/releases/download/v0.23.2/v6d-0.23.2.tar.gz"
	vineyard_cmd_tar := exec.Command("curl", "-L", vineyard_tar_url, "-o", "package.tar.gz")
	err := vineyard_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vineyard_zip_url := "https://github.com/v6d-io/v6d/releases/download/v0.23.2/v6d-0.23.2.zip"
	vineyard_cmd_zip := exec.Command("curl", "-L", vineyard_zip_url, "-o", "package.zip")
	err = vineyard_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vineyard_bin_url := "https://github.com/v6d-io/v6d/releases/download/v0.23.2/v6d-0.23.2.bin"
	vineyard_cmd_bin := exec.Command("curl", "-L", vineyard_bin_url, "-o", "binary.bin")
	err = vineyard_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vineyard_src_url := "https://github.com/v6d-io/v6d/releases/download/v0.23.2/v6d-0.23.2.src.tar.gz"
	vineyard_cmd_src := exec.Command("curl", "-L", vineyard_src_url, "-o", "source.tar.gz")
	err = vineyard_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vineyard_cmd_direct := exec.Command("./binary")
	err = vineyard_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: llvm")
	exec.Command("latte", "install", "llvm").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: apache-arrow")
	exec.Command("latte", "install", "apache-arrow").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: cpprestsdk")
	exec.Command("latte", "install", "cpprestsdk").Run()
	fmt.Println("Instalando dependencia: etcd")
	exec.Command("latte", "install", "etcd").Run()
	fmt.Println("Instalando dependencia: etcd-cpp-apiv3")
	exec.Command("latte", "install", "etcd-cpp-apiv3").Run()
	fmt.Println("Instalando dependencia: gflags")
	exec.Command("latte", "install", "gflags").Run()
	fmt.Println("Instalando dependencia: glog")
	exec.Command("latte", "install", "glog").Run()
	fmt.Println("Instalando dependencia: grpc")
	exec.Command("latte", "install", "grpc").Run()
	fmt.Println("Instalando dependencia: hiredis")
	exec.Command("latte", "install", "hiredis").Run()
	fmt.Println("Instalando dependencia: libgrape-lite")
	exec.Command("latte", "install", "libgrape-lite").Run()
	fmt.Println("Instalando dependencia: open-mpi")
	exec.Command("latte", "install", "open-mpi").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: protobuf")
	exec.Command("latte", "install", "protobuf").Run()
	fmt.Println("Instalando dependencia: redis")
	exec.Command("latte", "install", "redis").Run()
	fmt.Println("Instalando dependencia: abseil")
	exec.Command("latte", "install", "abseil").Run()
	fmt.Println("Instalando dependencia: c-ares")
	exec.Command("latte", "install", "c-ares").Run()
	fmt.Println("Instalando dependencia: re2")
	exec.Command("latte", "install", "re2").Run()
}
