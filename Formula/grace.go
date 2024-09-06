package main

// Grace - WYSIWYG 2D plotting tool for X11
// Homepage: https://plasma-gate.weizmann.ac.il/Grace/

import (
	"fmt"
	
	"os/exec"
)

func installGrace() {
	// Método 1: Descargar y extraer .tar.gz
	grace_tar_url := "https://deb.debian.org/debian/pool/main/g/grace/grace_5.1.25.orig.tar.gz"
	grace_cmd_tar := exec.Command("curl", "-L", grace_tar_url, "-o", "package.tar.gz")
	err := grace_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	grace_zip_url := "https://deb.debian.org/debian/pool/main/g/grace/grace_5.1.25.orig.zip"
	grace_cmd_zip := exec.Command("curl", "-L", grace_zip_url, "-o", "package.zip")
	err = grace_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	grace_bin_url := "https://deb.debian.org/debian/pool/main/g/grace/grace_5.1.25.orig.bin"
	grace_cmd_bin := exec.Command("curl", "-L", grace_bin_url, "-o", "binary.bin")
	err = grace_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	grace_src_url := "https://deb.debian.org/debian/pool/main/g/grace/grace_5.1.25.orig.src.tar.gz"
	grace_cmd_src := exec.Command("curl", "-L", grace_src_url, "-o", "source.tar.gz")
	err = grace_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	grace_cmd_direct := exec.Command("./binary")
	err = grace_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: fftw")
	exec.Command("latte", "install", "fftw").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libice")
	exec.Command("latte", "install", "libice").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libsm")
	exec.Command("latte", "install", "libsm").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxext")
	exec.Command("latte", "install", "libxext").Run()
	fmt.Println("Instalando dependencia: libxmu")
	exec.Command("latte", "install", "libxmu").Run()
	fmt.Println("Instalando dependencia: libxp")
	exec.Command("latte", "install", "libxp").Run()
	fmt.Println("Instalando dependencia: libxpm")
	exec.Command("latte", "install", "libxpm").Run()
	fmt.Println("Instalando dependencia: libxt")
	exec.Command("latte", "install", "libxt").Run()
	fmt.Println("Instalando dependencia: openmotif")
	exec.Command("latte", "install", "openmotif").Run()
}
