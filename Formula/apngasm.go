package main

// Apngasm - Next generation of apngasm, the APNG assembler
// Homepage: https://github.com/apngasm/apngasm

import (
	"fmt"
	
	"os/exec"
)

func installApngasm() {
	// Método 1: Descargar y extraer .tar.gz
	apngasm_tar_url := "https://github.com/apngasm/apngasm/archive/refs/tags/3.1.10.tar.gz"
	apngasm_cmd_tar := exec.Command("curl", "-L", apngasm_tar_url, "-o", "package.tar.gz")
	err := apngasm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	apngasm_zip_url := "https://github.com/apngasm/apngasm/archive/refs/tags/3.1.10.zip"
	apngasm_cmd_zip := exec.Command("curl", "-L", apngasm_zip_url, "-o", "package.zip")
	err = apngasm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	apngasm_bin_url := "https://github.com/apngasm/apngasm/archive/refs/tags/3.1.10.bin"
	apngasm_cmd_bin := exec.Command("curl", "-L", apngasm_bin_url, "-o", "binary.bin")
	err = apngasm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	apngasm_src_url := "https://github.com/apngasm/apngasm/archive/refs/tags/3.1.10.src.tar.gz"
	apngasm_cmd_src := exec.Command("curl", "-L", apngasm_src_url, "-o", "source.tar.gz")
	err = apngasm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	apngasm_cmd_direct := exec.Command("./binary")
	err = apngasm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: lzlib")
	exec.Command("latte", "install", "lzlib").Run()
}
