package main

// MsgpackCxx - MessagePack implementation for C++ / msgpack.org[C++]
// Homepage: https://msgpack.org/

import (
	"fmt"
	
	"os/exec"
)

func installMsgpackCxx() {
	// Método 1: Descargar y extraer .tar.gz
	msgpackcxx_tar_url := "https://github.com/msgpack/msgpack-c/releases/download/cpp-6.1.1/msgpack-cxx-6.1.1.tar.gz"
	msgpackcxx_cmd_tar := exec.Command("curl", "-L", msgpackcxx_tar_url, "-o", "package.tar.gz")
	err := msgpackcxx_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	msgpackcxx_zip_url := "https://github.com/msgpack/msgpack-c/releases/download/cpp-6.1.1/msgpack-cxx-6.1.1.zip"
	msgpackcxx_cmd_zip := exec.Command("curl", "-L", msgpackcxx_zip_url, "-o", "package.zip")
	err = msgpackcxx_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	msgpackcxx_bin_url := "https://github.com/msgpack/msgpack-c/releases/download/cpp-6.1.1/msgpack-cxx-6.1.1.bin"
	msgpackcxx_cmd_bin := exec.Command("curl", "-L", msgpackcxx_bin_url, "-o", "binary.bin")
	err = msgpackcxx_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	msgpackcxx_src_url := "https://github.com/msgpack/msgpack-c/releases/download/cpp-6.1.1/msgpack-cxx-6.1.1.src.tar.gz"
	msgpackcxx_cmd_src := exec.Command("curl", "-L", msgpackcxx_src_url, "-o", "source.tar.gz")
	err = msgpackcxx_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	msgpackcxx_cmd_direct := exec.Command("./binary")
	err = msgpackcxx_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
}
