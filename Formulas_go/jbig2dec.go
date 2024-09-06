package main

// Jbig2dec - JBIG2 decoder and library (for monochrome documents)
// Homepage: https://github.com/ArtifexSoftware/jbig2dec

import (
	"fmt"
	
	"os/exec"
)

func installJbig2dec() {
	// Método 1: Descargar y extraer .tar.gz
	jbig2dec_tar_url := "https://github.com/ArtifexSoftware/jbig2dec/archive/refs/tags/0.20.tar.gz"
	jbig2dec_cmd_tar := exec.Command("curl", "-L", jbig2dec_tar_url, "-o", "package.tar.gz")
	err := jbig2dec_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jbig2dec_zip_url := "https://github.com/ArtifexSoftware/jbig2dec/archive/refs/tags/0.20.zip"
	jbig2dec_cmd_zip := exec.Command("curl", "-L", jbig2dec_zip_url, "-o", "package.zip")
	err = jbig2dec_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jbig2dec_bin_url := "https://github.com/ArtifexSoftware/jbig2dec/archive/refs/tags/0.20.bin"
	jbig2dec_cmd_bin := exec.Command("curl", "-L", jbig2dec_bin_url, "-o", "binary.bin")
	err = jbig2dec_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jbig2dec_src_url := "https://github.com/ArtifexSoftware/jbig2dec/archive/refs/tags/0.20.src.tar.gz"
	jbig2dec_cmd_src := exec.Command("curl", "-L", jbig2dec_src_url, "-o", "source.tar.gz")
	err = jbig2dec_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jbig2dec_cmd_direct := exec.Command("./binary")
	err = jbig2dec_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
