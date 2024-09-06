package main

// AbiDumper - Dump ABI of an ELF object containing DWARF debug info
// Homepage: https://github.com/lvc/abi-dumper

import (
	"fmt"
	
	"os/exec"
)

func installAbiDumper() {
	// Método 1: Descargar y extraer .tar.gz
	abidumper_tar_url := "https://github.com/lvc/abi-dumper/archive/refs/tags/1.2.tar.gz"
	abidumper_cmd_tar := exec.Command("curl", "-L", abidumper_tar_url, "-o", "package.tar.gz")
	err := abidumper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	abidumper_zip_url := "https://github.com/lvc/abi-dumper/archive/refs/tags/1.2.zip"
	abidumper_cmd_zip := exec.Command("curl", "-L", abidumper_zip_url, "-o", "package.zip")
	err = abidumper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	abidumper_bin_url := "https://github.com/lvc/abi-dumper/archive/refs/tags/1.2.bin"
	abidumper_cmd_bin := exec.Command("curl", "-L", abidumper_bin_url, "-o", "binary.bin")
	err = abidumper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	abidumper_src_url := "https://github.com/lvc/abi-dumper/archive/refs/tags/1.2.src.tar.gz"
	abidumper_cmd_src := exec.Command("curl", "-L", abidumper_src_url, "-o", "source.tar.gz")
	err = abidumper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	abidumper_cmd_direct := exec.Command("./binary")
	err = abidumper_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: abi-compliance-checker")
exec.Command("latte", "install", "abi-compliance-checker")
	fmt.Println("Instalando dependencia: elfutils")
exec.Command("latte", "install", "elfutils")
	fmt.Println("Instalando dependencia: universal-ctags")
exec.Command("latte", "install", "universal-ctags")
	fmt.Println("Instalando dependencia: vtable-dumper")
exec.Command("latte", "install", "vtable-dumper")
}
