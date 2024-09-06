package main

// Jxrlib - Tools for JPEG-XR image encoding/decoding
// Homepage: https://tracker.debian.org/pkg/jxrlib

import (
	"fmt"
	
	"os/exec"
)

func installJxrlib() {
	// Método 1: Descargar y extraer .tar.gz
	jxrlib_tar_url := "https://deb.debian.org/debian/pool/main/j/jxrlib/jxrlib_1.2~git20170615.f752187.orig.tar.xz"
	jxrlib_cmd_tar := exec.Command("curl", "-L", jxrlib_tar_url, "-o", "package.tar.gz")
	err := jxrlib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jxrlib_zip_url := "https://deb.debian.org/debian/pool/main/j/jxrlib/jxrlib_1.2~git20170615.f752187.orig.tar.xz"
	jxrlib_cmd_zip := exec.Command("curl", "-L", jxrlib_zip_url, "-o", "package.zip")
	err = jxrlib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jxrlib_bin_url := "https://deb.debian.org/debian/pool/main/j/jxrlib/jxrlib_1.2~git20170615.f752187.orig.tar.xz"
	jxrlib_cmd_bin := exec.Command("curl", "-L", jxrlib_bin_url, "-o", "binary.bin")
	err = jxrlib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jxrlib_src_url := "https://deb.debian.org/debian/pool/main/j/jxrlib/jxrlib_1.2~git20170615.f752187.orig.tar.xz"
	jxrlib_cmd_src := exec.Command("curl", "-L", jxrlib_src_url, "-o", "source.tar.gz")
	err = jxrlib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jxrlib_cmd_direct := exec.Command("./binary")
	err = jxrlib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
}
