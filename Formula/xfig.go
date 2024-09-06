package main

// Xfig - Facility for interactive generation of figures
// Homepage: https://mcj.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installXfig() {
	// Método 1: Descargar y extraer .tar.gz
	xfig_tar_url := "https://downloads.sourceforge.net/mcj/xfig-3.2.9.tar.xz"
	xfig_cmd_tar := exec.Command("curl", "-L", xfig_tar_url, "-o", "package.tar.gz")
	err := xfig_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xfig_zip_url := "https://downloads.sourceforge.net/mcj/xfig-3.2.9.tar.xz"
	xfig_cmd_zip := exec.Command("curl", "-L", xfig_zip_url, "-o", "package.zip")
	err = xfig_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xfig_bin_url := "https://downloads.sourceforge.net/mcj/xfig-3.2.9.tar.xz"
	xfig_cmd_bin := exec.Command("curl", "-L", xfig_bin_url, "-o", "binary.bin")
	err = xfig_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xfig_src_url := "https://downloads.sourceforge.net/mcj/xfig-3.2.9.tar.xz"
	xfig_cmd_src := exec.Command("curl", "-L", xfig_src_url, "-o", "source.tar.gz")
	err = xfig_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xfig_cmd_direct := exec.Command("./binary")
	err = xfig_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: fig2dev")
	exec.Command("latte", "install", "fig2dev").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: ghostscript")
	exec.Command("latte", "install", "ghostscript").Run()
	fmt.Println("Instalando dependencia: jpeg-turbo")
	exec.Command("latte", "install", "jpeg-turbo").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: libtiff")
	exec.Command("latte", "install", "libtiff").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxaw3d")
	exec.Command("latte", "install", "libxaw3d").Run()
	fmt.Println("Instalando dependencia: libxft")
	exec.Command("latte", "install", "libxft").Run()
	fmt.Println("Instalando dependencia: libxpm")
	exec.Command("latte", "install", "libxpm").Run()
	fmt.Println("Instalando dependencia: libxt")
	exec.Command("latte", "install", "libxt").Run()
	fmt.Println("Instalando dependencia: gnu-sed")
	exec.Command("latte", "install", "gnu-sed").Run()
}
