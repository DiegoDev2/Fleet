package main

// Jbig2enc - JBIG2 encoder (for monochrome documents)
// Homepage: https://github.com/agl/jbig2enc

import (
	"fmt"
	
	"os/exec"
)

func installJbig2enc() {
	// Método 1: Descargar y extraer .tar.gz
	jbig2enc_tar_url := "https://github.com/agl/jbig2enc/archive/refs/tags/0.29.tar.gz"
	jbig2enc_cmd_tar := exec.Command("curl", "-L", jbig2enc_tar_url, "-o", "package.tar.gz")
	err := jbig2enc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jbig2enc_zip_url := "https://github.com/agl/jbig2enc/archive/refs/tags/0.29.zip"
	jbig2enc_cmd_zip := exec.Command("curl", "-L", jbig2enc_zip_url, "-o", "package.zip")
	err = jbig2enc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jbig2enc_bin_url := "https://github.com/agl/jbig2enc/archive/refs/tags/0.29.bin"
	jbig2enc_cmd_bin := exec.Command("curl", "-L", jbig2enc_bin_url, "-o", "binary.bin")
	err = jbig2enc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jbig2enc_src_url := "https://github.com/agl/jbig2enc/archive/refs/tags/0.29.src.tar.gz"
	jbig2enc_cmd_src := exec.Command("curl", "-L", jbig2enc_src_url, "-o", "source.tar.gz")
	err = jbig2enc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jbig2enc_cmd_direct := exec.Command("./binary")
	err = jbig2enc_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: leptonica")
exec.Command("latte", "install", "leptonica")
	fmt.Println("Instalando dependencia: giflib")
exec.Command("latte", "install", "giflib")
	fmt.Println("Instalando dependencia: jpeg-turbo")
exec.Command("latte", "install", "jpeg-turbo")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: libtiff")
exec.Command("latte", "install", "libtiff")
	fmt.Println("Instalando dependencia: webp")
exec.Command("latte", "install", "webp")
}
