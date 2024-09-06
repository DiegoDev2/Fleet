package main

// CmuSphinxbase - Lightweight speech recognition engine for mobile devices
// Homepage: https://cmusphinx.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installCmuSphinxbase() {
	// Método 1: Descargar y extraer .tar.gz
	cmusphinxbase_tar_url := "https://downloads.sourceforge.net/project/cmusphinx/sphinxbase/0.8/sphinxbase-0.8.tar.gz"
	cmusphinxbase_cmd_tar := exec.Command("curl", "-L", cmusphinxbase_tar_url, "-o", "package.tar.gz")
	err := cmusphinxbase_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cmusphinxbase_zip_url := "https://downloads.sourceforge.net/project/cmusphinx/sphinxbase/0.8/sphinxbase-0.8.zip"
	cmusphinxbase_cmd_zip := exec.Command("curl", "-L", cmusphinxbase_zip_url, "-o", "package.zip")
	err = cmusphinxbase_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cmusphinxbase_bin_url := "https://downloads.sourceforge.net/project/cmusphinx/sphinxbase/0.8/sphinxbase-0.8.bin"
	cmusphinxbase_cmd_bin := exec.Command("curl", "-L", cmusphinxbase_bin_url, "-o", "binary.bin")
	err = cmusphinxbase_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cmusphinxbase_src_url := "https://downloads.sourceforge.net/project/cmusphinx/sphinxbase/0.8/sphinxbase-0.8.src.tar.gz"
	cmusphinxbase_cmd_src := exec.Command("curl", "-L", cmusphinxbase_src_url, "-o", "source.tar.gz")
	err = cmusphinxbase_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cmusphinxbase_cmd_direct := exec.Command("./binary")
	err = cmusphinxbase_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: swig")
exec.Command("latte", "install", "swig")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libsamplerate")
exec.Command("latte", "install", "libsamplerate")
	fmt.Println("Instalando dependencia: libsndfile")
exec.Command("latte", "install", "libsndfile")
}
