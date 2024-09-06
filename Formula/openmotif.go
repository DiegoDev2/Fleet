package main

// Openmotif - LGPL release of the Motif toolkit
// Homepage: https://motif.ics.com/motif

import (
	"fmt"
	
	"os/exec"
)

func installOpenmotif() {
	// Método 1: Descargar y extraer .tar.gz
	openmotif_tar_url := "https://downloads.sourceforge.net/project/motif/Motif%202.3.8%20Source%20Code/motif-2.3.8.tar.gz"
	openmotif_cmd_tar := exec.Command("curl", "-L", openmotif_tar_url, "-o", "package.tar.gz")
	err := openmotif_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openmotif_zip_url := "https://downloads.sourceforge.net/project/motif/Motif%202.3.8%20Source%20Code/motif-2.3.8.zip"
	openmotif_cmd_zip := exec.Command("curl", "-L", openmotif_zip_url, "-o", "package.zip")
	err = openmotif_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openmotif_bin_url := "https://downloads.sourceforge.net/project/motif/Motif%202.3.8%20Source%20Code/motif-2.3.8.bin"
	openmotif_cmd_bin := exec.Command("curl", "-L", openmotif_bin_url, "-o", "binary.bin")
	err = openmotif_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openmotif_src_url := "https://downloads.sourceforge.net/project/motif/Motif%202.3.8%20Source%20Code/motif-2.3.8.src.tar.gz"
	openmotif_cmd_src := exec.Command("curl", "-L", openmotif_src_url, "-o", "source.tar.gz")
	err = openmotif_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openmotif_cmd_direct := exec.Command("./binary")
	err = openmotif_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
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
	fmt.Println("Instalando dependencia: libxft")
	exec.Command("latte", "install", "libxft").Run()
	fmt.Println("Instalando dependencia: libxmu")
	exec.Command("latte", "install", "libxmu").Run()
	fmt.Println("Instalando dependencia: libxp")
	exec.Command("latte", "install", "libxp").Run()
	fmt.Println("Instalando dependencia: libxt")
	exec.Command("latte", "install", "libxt").Run()
	fmt.Println("Instalando dependencia: xbitmaps")
	exec.Command("latte", "install", "xbitmaps").Run()
}
