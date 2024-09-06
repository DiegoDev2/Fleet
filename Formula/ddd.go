package main

// Ddd - Graphical front-end for command-line debuggers
// Homepage: https://www.gnu.org/software/ddd/

import (
	"fmt"
	
	"os/exec"
)

func installDdd() {
	// Método 1: Descargar y extraer .tar.gz
	ddd_tar_url := "https://ftp.gnu.org/gnu/ddd/ddd-3.4.1.tar.gz"
	ddd_cmd_tar := exec.Command("curl", "-L", ddd_tar_url, "-o", "package.tar.gz")
	err := ddd_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ddd_zip_url := "https://ftp.gnu.org/gnu/ddd/ddd-3.4.1.zip"
	ddd_cmd_zip := exec.Command("curl", "-L", ddd_zip_url, "-o", "package.zip")
	err = ddd_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ddd_bin_url := "https://ftp.gnu.org/gnu/ddd/ddd-3.4.1.bin"
	ddd_cmd_bin := exec.Command("curl", "-L", ddd_bin_url, "-o", "binary.bin")
	err = ddd_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ddd_src_url := "https://ftp.gnu.org/gnu/ddd/ddd-3.4.1.src.tar.gz"
	ddd_cmd_src := exec.Command("curl", "-L", ddd_src_url, "-o", "source.tar.gz")
	err = ddd_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ddd_cmd_direct := exec.Command("./binary")
	err = ddd_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxaw")
	exec.Command("latte", "install", "libxaw").Run()
	fmt.Println("Instalando dependencia: libxft")
	exec.Command("latte", "install", "libxft").Run()
	fmt.Println("Instalando dependencia: libxmu")
	exec.Command("latte", "install", "libxmu").Run()
	fmt.Println("Instalando dependencia: libxpm")
	exec.Command("latte", "install", "libxpm").Run()
	fmt.Println("Instalando dependencia: libxt")
	exec.Command("latte", "install", "libxt").Run()
	fmt.Println("Instalando dependencia: openmotif")
	exec.Command("latte", "install", "openmotif").Run()
	fmt.Println("Instalando dependencia: gnu-sed")
	exec.Command("latte", "install", "gnu-sed").Run()
	fmt.Println("Instalando dependencia: libice")
	exec.Command("latte", "install", "libice").Run()
	fmt.Println("Instalando dependencia: libsm")
	exec.Command("latte", "install", "libsm").Run()
	fmt.Println("Instalando dependencia: libxext")
	exec.Command("latte", "install", "libxext").Run()
	fmt.Println("Instalando dependencia: libxp")
	exec.Command("latte", "install", "libxp").Run()
	fmt.Println("Instalando dependencia: gdb")
	exec.Command("latte", "install", "gdb").Run()
	fmt.Println("Instalando dependencia: gdb")
	exec.Command("latte", "install", "gdb").Run()
}
