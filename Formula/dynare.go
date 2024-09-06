package main

// Dynare - Platform for economic models, particularly DSGE and OLG models
// Homepage: https://www.dynare.org/

import (
	"fmt"
	
	"os/exec"
)

func installDynare() {
	// Método 1: Descargar y extraer .tar.gz
	dynare_tar_url := "https://www.dynare.org/release/source/dynare-6.1.tar.xz"
	dynare_cmd_tar := exec.Command("curl", "-L", dynare_tar_url, "-o", "package.tar.gz")
	err := dynare_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dynare_zip_url := "https://www.dynare.org/release/source/dynare-6.1.tar.xz"
	dynare_cmd_zip := exec.Command("curl", "-L", dynare_zip_url, "-o", "package.zip")
	err = dynare_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dynare_bin_url := "https://www.dynare.org/release/source/dynare-6.1.tar.xz"
	dynare_cmd_bin := exec.Command("curl", "-L", dynare_bin_url, "-o", "binary.bin")
	err = dynare_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dynare_src_url := "https://www.dynare.org/release/source/dynare-6.1.tar.xz"
	dynare_cmd_src := exec.Command("curl", "-L", dynare_src_url, "-o", "source.tar.gz")
	err = dynare_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dynare_cmd_direct := exec.Command("./binary")
	err = dynare_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
	exec.Command("latte", "install", "bison").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: cweb")
	exec.Command("latte", "install", "cweb").Run()
	fmt.Println("Instalando dependencia: flex")
	exec.Command("latte", "install", "flex").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: fftw")
	exec.Command("latte", "install", "fftw").Run()
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
	fmt.Println("Instalando dependencia: gsl")
	exec.Command("latte", "install", "gsl").Run()
	fmt.Println("Instalando dependencia: hdf5")
	exec.Command("latte", "install", "hdf5").Run()
	fmt.Println("Instalando dependencia: libmatio")
	exec.Command("latte", "install", "libmatio").Run()
	fmt.Println("Instalando dependencia: metis")
	exec.Command("latte", "install", "metis").Run()
	fmt.Println("Instalando dependencia: octave")
	exec.Command("latte", "install", "octave").Run()
	fmt.Println("Instalando dependencia: openblas")
	exec.Command("latte", "install", "openblas").Run()
	fmt.Println("Instalando dependencia: suite-sparse")
	exec.Command("latte", "install", "suite-sparse").Run()
}
