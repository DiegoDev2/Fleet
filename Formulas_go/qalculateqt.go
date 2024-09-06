package main

// QalculateQt - Multi-purpose desktop calculator
// Homepage: https://qalculate.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installQalculateQt() {
	// Método 1: Descargar y extraer .tar.gz
	qalculateqt_tar_url := "https://github.com/Qalculate/qalculate-qt/releases/download/v5.2.0/qalculate-qt-5.2.0.tar.gz"
	qalculateqt_cmd_tar := exec.Command("curl", "-L", qalculateqt_tar_url, "-o", "package.tar.gz")
	err := qalculateqt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	qalculateqt_zip_url := "https://github.com/Qalculate/qalculate-qt/releases/download/v5.2.0/qalculate-qt-5.2.0.zip"
	qalculateqt_cmd_zip := exec.Command("curl", "-L", qalculateqt_zip_url, "-o", "package.zip")
	err = qalculateqt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	qalculateqt_bin_url := "https://github.com/Qalculate/qalculate-qt/releases/download/v5.2.0/qalculate-qt-5.2.0.bin"
	qalculateqt_cmd_bin := exec.Command("curl", "-L", qalculateqt_bin_url, "-o", "binary.bin")
	err = qalculateqt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	qalculateqt_src_url := "https://github.com/Qalculate/qalculate-qt/releases/download/v5.2.0/qalculate-qt-5.2.0.src.tar.gz"
	qalculateqt_cmd_src := exec.Command("curl", "-L", qalculateqt_src_url, "-o", "source.tar.gz")
	err = qalculateqt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	qalculateqt_cmd_direct := exec.Command("./binary")
	err = qalculateqt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libqalculate")
exec.Command("latte", "install", "libqalculate")
	fmt.Println("Instalando dependencia: qt")
exec.Command("latte", "install", "qt")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: mpfr")
exec.Command("latte", "install", "mpfr")
}
