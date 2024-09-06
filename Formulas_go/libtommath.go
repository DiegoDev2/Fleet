package main

// Libtommath - C library for number theoretic multiple-precision integers
// Homepage: https://www.libtom.net/LibTomMath/

import (
	"fmt"
	
	"os/exec"
)

func installLibtommath() {
	// Método 1: Descargar y extraer .tar.gz
	libtommath_tar_url := "https://github.com/libtom/libtommath/releases/download/v1.3.0/ltm-1.3.0.tar.xz"
	libtommath_cmd_tar := exec.Command("curl", "-L", libtommath_tar_url, "-o", "package.tar.gz")
	err := libtommath_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libtommath_zip_url := "https://github.com/libtom/libtommath/releases/download/v1.3.0/ltm-1.3.0.tar.xz"
	libtommath_cmd_zip := exec.Command("curl", "-L", libtommath_zip_url, "-o", "package.zip")
	err = libtommath_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libtommath_bin_url := "https://github.com/libtom/libtommath/releases/download/v1.3.0/ltm-1.3.0.tar.xz"
	libtommath_cmd_bin := exec.Command("curl", "-L", libtommath_bin_url, "-o", "binary.bin")
	err = libtommath_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libtommath_src_url := "https://github.com/libtom/libtommath/releases/download/v1.3.0/ltm-1.3.0.tar.xz"
	libtommath_cmd_src := exec.Command("curl", "-L", libtommath_src_url, "-o", "source.tar.gz")
	err = libtommath_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libtommath_cmd_direct := exec.Command("./binary")
	err = libtommath_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}
