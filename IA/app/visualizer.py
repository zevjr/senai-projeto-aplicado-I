import matplotlib.pyplot as plt
from typing import Dict, Optional
from pathlib import Path


class IncidentVisualizer:
    """Gera visualizações a partir dos dados de incidentes"""

    def __init__(self, risk_distribution: Dict[int, int]):
        self.risk_distribution = risk_distribution
        plt.style.use('ggplot')

    def plot_risk_distribution(self, save_path: Optional[str] = None) -> None:
        """Gera gráfico de distribuição de risco"""
        if not self.risk_distribution:
            print("Nenhum dado disponível para plotagem")
            return

        levels = sorted(self.risk_distribution.keys())
        counts = [self.risk_distribution[l] for l in levels]

        plt.figure(figsize=(10, 6))
        bars = plt.bar(levels, counts, color='#4e79a7')

        # Adiciona rótulos
        for bar in bars:
            height = bar.get_height()
            plt.text(bar.get_x() + bar.get_width() / 2., height,
                     f'{int(height)}', ha='center', va='bottom')

        plt.title('Distribuição de Níveis de Risco', pad=20)
        plt.xlabel('Nível de Risco', labelpad=10)
        plt.ylabel('Número de Ocorrências', labelpad=10)
        plt.xticks(range(1, 11))
        plt.grid(axis='y', alpha=0.3)
        plt.tight_layout()

        if save_path:
            Path(save_path).parent.mkdir(parents=True, exist_ok=True)
            plt.savefig(save_path, dpi=300)
            plt.close()
            print(f"Gráfico salvo em {save_path}")
        else:
            plt.show()