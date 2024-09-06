package main

// EtcdCppApiv3 - C++ implementation for etcd's v3 client API, i.e., ETCDCTL_API=3
// Homepage: https://github.com/etcd-cpp-apiv3/etcd-cpp-apiv3

import (
	"fmt"
	
	"os/exec"
)

func installEtcdCppApiv3() {
	// Método 1: Descargar y extraer .tar.gz
	etcdcppapiv3_tar_url := "https://github.com/etcd-cpp-apiv3/etcd-cpp-apiv3/archive/refs/tags/v0.15.4.tar.gz"
	etcdcppapiv3_cmd_tar := exec.Command("curl", "-L", etcdcppapiv3_tar_url, "-o", "package.tar.gz")
	err := etcdcppapiv3_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	etcdcppapiv3_zip_url := "https://github.com/etcd-cpp-apiv3/etcd-cpp-apiv3/archive/refs/tags/v0.15.4.zip"
	etcdcppapiv3_cmd_zip := exec.Command("curl", "-L", etcdcppapiv3_zip_url, "-o", "package.zip")
	err = etcdcppapiv3_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	etcdcppapiv3_bin_url := "https://github.com/etcd-cpp-apiv3/etcd-cpp-apiv3/archive/refs/tags/v0.15.4.bin"
	etcdcppapiv3_cmd_bin := exec.Command("curl", "-L", etcdcppapiv3_bin_url, "-o", "binary.bin")
	err = etcdcppapiv3_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	etcdcppapiv3_src_url := "https://github.com/etcd-cpp-apiv3/etcd-cpp-apiv3/archive/refs/tags/v0.15.4.src.tar.gz"
	etcdcppapiv3_cmd_src := exec.Command("curl", "-L", etcdcppapiv3_src_url, "-o", "source.tar.gz")
	err = etcdcppapiv3_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	etcdcppapiv3_cmd_direct := exec.Command("./binary")
	err = etcdcppapiv3_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: etcd")
	exec.Command("latte", "install", "etcd").Run()
	fmt.Println("Instalando dependencia: abseil")
	exec.Command("latte", "install", "abseil").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: c-ares")
	exec.Command("latte", "install", "c-ares").Run()
	fmt.Println("Instalando dependencia: cpprestsdk")
	exec.Command("latte", "install", "cpprestsdk").Run()
	fmt.Println("Instalando dependencia: grpc")
	exec.Command("latte", "install", "grpc").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: protobuf")
	exec.Command("latte", "install", "protobuf").Run()
	fmt.Println("Instalando dependencia: re2")
	exec.Command("latte", "install", "re2").Run()
}
