package main

// Valgrind - Dynamic analysis tools (memory, debug, profiling)
// Homepage: https://www.valgrind.org/

import (
	"fmt"
	
	"os/exec"
)

func installValgrind() {
	// Método 1: Descargar y extraer .tar.gz
	valgrind_tar_url := "https://sourceware.org/pub/valgrind/valgrind-3.23.0.tar.bz2"
	valgrind_cmd_tar := exec.Command("curl", "-L", valgrind_tar_url, "-o", "package.tar.gz")
	err := valgrind_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	valgrind_zip_url := "https://sourceware.org/pub/valgrind/valgrind-3.23.0.tar.bz2"
	valgrind_cmd_zip := exec.Command("curl", "-L", valgrind_zip_url, "-o", "package.zip")
	err = valgrind_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	valgrind_bin_url := "https://sourceware.org/pub/valgrind/valgrind-3.23.0.tar.bz2"
	valgrind_cmd_bin := exec.Command("curl", "-L", valgrind_bin_url, "-o", "binary.bin")
	err = valgrind_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	valgrind_src_url := "https://sourceware.org/pub/valgrind/valgrind-3.23.0.tar.bz2"
	valgrind_cmd_src := exec.Command("curl", "-L", valgrind_src_url, "-o", "source.tar.gz")
	err = valgrind_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	valgrind_cmd_direct := exec.Command("./binary")
	err = valgrind_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
