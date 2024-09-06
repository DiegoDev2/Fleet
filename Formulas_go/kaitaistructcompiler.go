package main

// KaitaiStructCompiler - Compiler for generating binary data parsers
// Homepage: https://kaitai.io/

import (
	"fmt"
	
	"os/exec"
)

func installKaitaiStructCompiler() {
	// Método 1: Descargar y extraer .tar.gz
	kaitaistructcompiler_tar_url := "https://github.com/kaitai-io/kaitai_struct_compiler/releases/download/0.10/kaitai-struct-compiler-0.10.zip"
	kaitaistructcompiler_cmd_tar := exec.Command("curl", "-L", kaitaistructcompiler_tar_url, "-o", "package.tar.gz")
	err := kaitaistructcompiler_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	kaitaistructcompiler_zip_url := "https://github.com/kaitai-io/kaitai_struct_compiler/releases/download/0.10/kaitai-struct-compiler-0.10.zip"
	kaitaistructcompiler_cmd_zip := exec.Command("curl", "-L", kaitaistructcompiler_zip_url, "-o", "package.zip")
	err = kaitaistructcompiler_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	kaitaistructcompiler_bin_url := "https://github.com/kaitai-io/kaitai_struct_compiler/releases/download/0.10/kaitai-struct-compiler-0.10.zip"
	kaitaistructcompiler_cmd_bin := exec.Command("curl", "-L", kaitaistructcompiler_bin_url, "-o", "binary.bin")
	err = kaitaistructcompiler_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	kaitaistructcompiler_src_url := "https://github.com/kaitai-io/kaitai_struct_compiler/releases/download/0.10/kaitai-struct-compiler-0.10.zip"
	kaitaistructcompiler_cmd_src := exec.Command("curl", "-L", kaitaistructcompiler_src_url, "-o", "source.tar.gz")
	err = kaitaistructcompiler_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	kaitaistructcompiler_cmd_direct := exec.Command("./binary")
	err = kaitaistructcompiler_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}
